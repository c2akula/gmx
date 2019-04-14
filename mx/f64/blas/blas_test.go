package blas

import (
	"math"
	"math/rand"
	"testing"
)

var x, y []float64

func init() {
	x = make([]float64, 514)
	y = make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		x[i] = rand.Float64()
		y[i] = rand.Float64()
	}
}

func test_asum(n int, x []float64, incx int) float64 {
	s := 0.
	for k := 0; n > 0; k += incx {
		n--
		s += math.Abs(x[k])
	}
	return s
}

func TestASum(t *testing.T) {
	x := []float64{-1, -2, 3, -4, -5, -6, 7, -8, 9, 1, -2, 3, -4, 5, -6, -7, 8}
	for incx := 1; incx < 9; incx++ {
		n := len(x) / incx
		got := ASum(n, x, incx)
		exp := test_asum(n, x, incx)
		if got != exp {
			t.Fatalf("test: [%d] failed!. inc:%d got:%f, exp:%f\n", incx-1, incx, got, exp)
		}
	}
}

func BenchmarkASum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ASum(len(x), x, 1)
	}
}

func test_axpy(n int, s float64, x []float64, incx int, y []float64, incy int) {
	for i, j := 0, 0; n > 0; {
		n--
		y[j] += s * x[i]
		i += incx
		j += incy
	}
}

func isequal(a, b []float64) (i int, ai, bi float64, eq bool) {
	for i := 0; i < len(x); i++ {
		if a[i] != b[i] {
			return i, a[i], b[i], false
		}
	}
	return 0, 0, 0, true
}

func TestAxpy(t *testing.T) {
	for _, s := range []float64{0, -1, 1, 3} {
		for inc := 1; inc < 9; inc++ {
			n := len(x) / inc
			r := make([]float64, len(x))
			e := make([]float64, len(x))
			copy(r, x)
			copy(e, x)
			Axpy(n, s, x, inc, r, inc)
			test_axpy(n, s, x, inc, e, inc)
			for i := 0; i < len(x); i++ {
				if r[i] != e[i] {
					t.Fatalf("s:%f inc:%d n:%d i:%d r:%f e:%f", s, inc, n, i, r[i], e[i])
				}
			}
		}
	}
}

func BenchmarkAxpy(b *testing.B) {
	y := make([]float64, len(x))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Axpy(len(x), -1, x, 1, y, 1)
	}
}

func axpy(n int, s float64, x []float64, incx int, y []float64, incy int) {
	i, j := 0, 0
	switch s {
	case 0:
	case 1:
		for ; n >= 4; n -= 4 {
			y[j] += x[i]
			i += incx
			j += incy
			y[j] += x[i]
			i += incx
			j += incy

			y[j] += x[i]
			i += incx
			j += incy
			y[j] += x[i]
			i += incx
			j += incy
		}

		for n != 0 {
			y[j] += x[i]
			i += incx
			j += incy

			n--
		}
	case -1:
		for ; n >= 4; n -= 4 {
			y[j] -= x[i]
			i += incx
			j += incy
			y[j] -= x[i]
			i += incx
			j += incy
			y[j] -= x[i]
			i += incx
			j += incy
			y[j] -= x[i]
			i += incx
			j += incy
		}
		for n != 0 {
			y[j] -= x[i]
			i += incx
			j += incy
			n--
		}
	default:
		for ; n >= 4; n -= 4 {
			y[j] += s * x[i]
			i += incx
			j += incy
			y[j] += s * x[i]
			i += incx
			j += incy
			y[j] += s * x[i]
			i += incx
			j += incy
			y[j] += s * x[i]
			i += incx
			j += incy
		}
		for n != 0 {
			y[j] += s * x[i]
			i += incx
			j += incy
			n--
		}
	}
}

func TestAxpy2(t *testing.T) {
	x := make([]float64, 20)
	for i := range x {
		x[i] = float64(rand.Intn(len(x)))
	}
	r := make([]float64, len(x))
	copy(r, x)
	e := []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9}
	axpy(len(x), -1, x, 1, r, 1)
	for i := 0; i < len(x); i++ {
		if e[i] != r[i] {
			t.Fatalf("failed!. got: %v, exp: %v\n", r, e)
		}
	}
}
func BenchmarkAxpy2(b *testing.B) {
	y := make([]float64, len(x))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		axpy(len(x), -1, x, 1, y, 1)
	}
}

func test_copy(n int, x []float64, incx int, y []float64, incy int) {
	for i, j := 0, 0; n > 0; n-- {
		y[j] = x[i]
		i += incx
		j += incy
	}
}

func TestCopy(t *testing.T) {
	r := make([]float64, len(x))
	e := make([]float64, len(x))
	for inc := 1; inc < 9; inc++ {
		n := len(x) / inc
		Copy(n, x, inc, r, inc)
		test_copy(n, x, inc, e, inc)
		for i := 0; i < len(x); i++ {
			if r[i] != e[i] {
				t.Fatalf("inc:%d n:%d i:%d r:%f e:%f", inc, n, i, r[i], e[i])
			}
		}
	}
}

func BenchmarkCopy_Inc(b *testing.B) {
	inc := 2
	n := len(x) / 2
	for i := 0; i < b.N; i++ {
		Copy(n, x, inc, y, inc)
	}
}
func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(len(x), x, 1, y, 1)
	}
}

func test_dot(n int, x []float64, incx int, y []float64, incy int) float64 {
	s := 0.
	for i, j := 0, 0; n > 0; n-- {
		s += x[i] * y[j]
		i += incx
		j += incy
	}
	return s
}

func TestDot(t *testing.T) {
	for inc := 1; inc < 9; inc++ {
		n := len(x) / inc
		r := Dot(n, x, inc, y, inc)
		e := test_dot(n, x, inc, y, inc)
		for i := 0; i < len(x); i++ {
			if r != e {
				t.Fatalf("inc:%d n:%d r:%f e:%f", inc, n, r, e)
			}
		}
	}
}

func BenchmarkDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dot(len(x), x, 1, y, 1)
	}
}

func test_nrm2(n int, x []float64, incx int) float64 {
	nrm := 0.0
	for i := 0; n > 0; n-- {
		nrm += x[i] * x[i]
		i += incx
	}
	return math.Sqrt(nrm)
}

func TestNrm2(t *testing.T) {
	for inc := 1; inc < 9; inc++ {
		n := len(x) / inc
		r := Nrm2(n, x, inc)
		e := test_nrm2(n, x, inc)
		for i := 0; i < len(x); i++ {
			if r != e {
				t.Fatalf("inc:%d n:%d r:%f e:%f", inc, n, r, e)
			}
		}
	}
}

func BenchmarkNrm2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nrm2(len(x), x, 1)
	}
}

func test_fill(n int, x []float64, incx int, s float64) {
	for i := 0; n > 0; n-- {
		x[i] = s
		i += incx
	}
}

func TestFill(t *testing.T) {
	e := make([]float64, len(x))
	r := make([]float64, len(x))

	for inc := 1; inc < 9; inc++ {
		copy(e, x)
		copy(r, x)
		n := len(x) / inc
		Fill(n, r, inc, -1)
		test_fill(n, e, inc, -1)
		for i := 0; i < len(x); i++ {
			if r[i] != e[i] {
				t.Fatalf("inc:%d n:%d i:%d r:%f e:%f", inc, n, i, r[i], e[i])
			}
		}
	}
}

func BenchmarkFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fill(len(x), x, 1, -1)
	}
}

func test_scale(n int, x []float64, incx int, s float64) {
	for i := 0; n > 0; n-- {
		x[i] *= s
		i += incx
	}
}

func TestScale(t *testing.T) {
	e := make([]float64, len(x))
	r := make([]float64, len(x))
	for inc := 1; inc < 9; inc++ {
		n := len(x) / inc
		copy(e, x)
		copy(r, x)
		test_scale(n, e, inc, -1)
		Scale(n, r, inc, -1)
		for i := 0; i < len(x); i++ {
			if r[i] != e[i] {
				t.Fatalf("inc:%d n:%d i:%d r:%f e:%f", inc, n, i, r[i], e[i])
			}
		}
	}
}

func BenchmarkScale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Scale(len(x), x, 1, -1)
	}
}

func test_swap(n int, x []float64, incx int, y []float64, incy int) {
	for i, j := 0, 0; n > 0; n-- {
		x[i], y[j] = y[j], x[i]
		i += incx
		j += incy
	}
}

func TestSwap(t *testing.T) {
	xr := make([]float64, len(x))
	yr := make([]float64, len(y))
	xe := make([]float64, len(x))
	ye := make([]float64, len(y))
	for inc := 1; inc < 9; inc++ {
		n := len(x) / inc
		copy(xr, x)
		copy(yr, y)
		copy(xe, x)
		copy(ye, y)

		Swap(n, xr, inc, yr, inc)
		test_swap(n, xe, inc, ye, inc)

		for i := 0; i < len(x); i++ {
			if xr[i] != xe[i] || yr[i] != ye[i] {
				t.Fatalf("inc:%d n:%d i:%d xr:%f xe:%f", inc, n, i, xr[i], xe[i])
				t.Fatalf("inc:%d n:%d i:%d yr:%f ye:%f", inc, n, i, yr[i], ye[i])
			}
		}
	}
}

func BenchmarkSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Swap(len(x), x, 1, y, 1)
	}
}

func test_max(n int, x []float64, incx int) (float64, int) {
	s := 0.0
	idx := 0
	for i := 0; i < len(x); i += incx {
		if x[i] > s {
			s = x[i]
			idx = i
		}
	}
	return s, idx
}

func test_amax(n int, x []float64, incx int) (float64, int) {
	s := 0.0
	idx := 0
	for i := 0; i < len(x); i += incx {
		xi := math.Abs(x[i])
		if xi > s {
			s = xi
			idx = i
		}
	}
	return s, idx
}

func TestMax(t *testing.T) {
	rmax, ri := Max(len(x), x, 1)
	emax, ei := test_max(len(x), x, 1)
	if rmax != emax || ri != ei {
		t.Fatalf("failed. rmax:%f, emax:%f, ri:%d, ei:%d\n", rmax, emax, ri, ei)
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(len(x), x, 1)
	}
}

func TestAMax(t *testing.T) {
	rmax, ri := AMax(len(x), x, 1)
	emax, ei := test_amax(len(x), x, 1)
	if rmax != emax || ri != ei {
		t.Fatalf("failed. rmax:%f, emax:%f, ri:%d, ei:%d\n", rmax, emax, ri, ei)
	}
}

func BenchmarkAMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AMax(len(x), x, 1)
	}
}
