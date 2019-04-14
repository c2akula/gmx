package blas

func ASum(n int, x []float64, incx int) (s float64) {
	for k := 0; n > 0; k += incx {
		n--
		xk := x[k]
		if xk < 0 {
			xk = -xk
		}
		s += xk
	}

	return
}
