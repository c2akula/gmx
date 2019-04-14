package blas

func Axpy(n int, s float64, x []float64, incx int, y []float64, incy int) {
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
