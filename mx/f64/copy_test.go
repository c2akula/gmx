package f64

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCopy(t *testing.T) {
	const (
		m, n = 4, 4
	)
	src := randmx(m, n)
	tests := []struct {
		dst, src, exp *Mx
	}{
		{
			// NT.NT
			Zeros(m, n),
			src.Dup(),
			src.Dup(),
		},
		{
			// T.NT
			Zeros(m, n),
			src.Dup().Transpose(),
			src.Dup().Transpose(),
		},
		{
			// 	NT.T
			Zeros(m, n).Transpose(),
			src.Dup(),
			src.Dup(),
		},
		{
			// 	T.T
			Zeros(m, n).Transpose(),
			src.Dup().Transpose(),
			src.Dup().Transpose(),
		},
	}

	for ti, tt := range tests {
		Copy(tt.dst, tt.src)
		if !IsEqual(tt.dst, tt.src) {
			t.Fatalf("test[%d] failed.\ndst: %sexpected: %s\n", ti, tt.dst, tt.exp)
		}
	}
}

func TestCopySlice(t *testing.T) {
	got := New([]float64{
		1, 2, 3, 1, 2, 3,
		4, 5, 6, 4, 5, 6,
		7, 8, 9, 7, 8, 9,
		10, 11, 12, 10, 11, 12,
	}, 4, 6)
	exp := New([]float64{
		4, 5, 6, 1, 2, 3,
		7, 8, 9, 4, 5, 6,
		7, 8, 9, 7, 8, 9,
		10, 11, 12, 10, 11, 12,
	}, 4, 6)
	fmt.Println("copying...")
	Copy(got.Blk(0, 0, 2, 3), got.Blk(1, 3, 2, 3))
	if !IsEqual(exp, got) {
		t.Fatalf("got=%sexp=%s\n", got, exp)
	}
}

func randmx(nr, nc int) *Mx {
	m := Zeros(nr, nc)
	for i := 0; i < m.len; i++ {
		m.Set(i, float64(rand.Intn(nr*nc)))
	}
	return m
}
