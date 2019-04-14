package blas

func Copy(n int, x []float64, incx int, y []float64, incy int) {
	if incx == 1 && incy == 1 {
		copy(y[:n], x[:n])
		return
	}

	i, j := 0, 0
	for ; n >= 4; n -= 4 {
		y[j] = x[i]
		i += incx
		j += incy
		y[j] = x[i]
		i += incx
		j += incy
		y[j] = x[i]
		i += incx
		j += incy
		y[j] = x[i]
		i += incx
		j += incy
	}

	for n != 0 {
		y[j] = x[i]
		i += incx
		j += incy
		n--
	}
}
