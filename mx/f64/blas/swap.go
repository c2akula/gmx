package blas

func Swap(n int, x []float64, incx int, y []float64, incy int) {
	i, j := 0, 0
	for ; n >= 4; n -= 4 {
		x[i], y[j] = y[j], x[i]
		i += incx
		j += incy
		x[i], y[j] = y[j], x[i]
		i += incx
		j += incy
		x[i], y[j] = y[j], x[i]
		i += incx
		j += incy
		x[i], y[j] = y[j], x[i]
		i += incx
		j += incy
	}

	for ; n != 0; n-- {
		x[i], y[j] = y[j], x[i]
		i += incx
		j += incy
	}
}
