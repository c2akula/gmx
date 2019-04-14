package blas

func Dot(n int, x []float64, incx int, y []float64, incy int) (s float64) {
	i, j := 0, 0
	for ; n >= 4; n -= 4 {
		s += x[i] * y[j]
		i += incx
		j += incy
		s += x[i] * y[j]
		i += incx
		j += incy
		s += x[i] * y[j]
		i += incx
		j += incy
		s += x[i] * y[j]
		i += incx
		j += incy
	}

	for n != 0 {
		s += x[i] * y[j]
		i += incx
		j += incy
		n--
	}
	return
}
