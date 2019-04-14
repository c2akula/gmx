package blas

func Scale(n int, x []float64, incx int, s float64) {
	i := 0
	for ; n >= 4; n -= 4 {
		x[i] *= s
		i += incx
		x[i] *= s
		i += incx
		x[i] *= s
		i += incx
		x[i] *= s
		i += incx
	}

	for ; n != 0; n-- {
		x[i] *= s
		i += incx
	}
}
