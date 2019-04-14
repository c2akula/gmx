package blas

func Max(n int, x []float64, incx int) (s float64, idx int) {
	s = x[0]
	for i := 1; n > 1; n-- {
		if x[i] > s {
			s = x[i]
			idx = i
		}
		i += incx
	}

	return
}

func AMax(n int, x []float64, incx int) (s float64, idx int) {
	s = x[0]
	if s < 0 {
		s = -s
	}
	for i := 1; n > 1; n-- {
		xi := x[i]
		if xi < 0 {
			xi = -xi
		}
		if xi > s {
			s = xi
			idx = i
		}
		i += incx
	}
	return
}
