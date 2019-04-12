package blas

import "math"

func Scale(n int, x []float64, incx int, s float64) {
	for i := 0; i < n; i += incx {
		x[i] *= s
	}
}

func Iaxpy(n int, a float64, x []float64, incx int, y []float64, incy int) {
	for i, ix, iy := 0, 0, 0; i < n; i, ix, iy = i+1, ix+incx, iy+incy {
		y[iy] += a * x[ix]
	}
}

func Uaxpy(n int, a float64, x, y []float64) {
	for i, xe := range x[:n] {
		y[i] += a * xe
	}
}

func IDot(n int, x []float64, incx int, y []float64, incy int) (s float64) {
	for i, ix, iy := 0, 0, 0; i < n; i, ix, iy = i+1, ix+incx, iy+incy {
		s += x[ix] * y[iy]
	}
	return
}

func UDot(n int, x, y []float64) (s float64) {
	for i, xe := range x[:n] {
		s += xe * y[i]
	}
	return
}

func Norm(n int, x []float64, incx int) (s float64) {
	for i, k := 0, 0; i < n; i, k = i+1, k+incx {
		xk := x[k]
		s += xk * xk
	}
	return math.Sqrt(s)
}
