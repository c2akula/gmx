package blas

func Fill(n int, x []float64, incx int, s float64) {
	i := 0
	for ; n >= 4; n -= 4 {
		x[i] = s
		i += incx

		x[i] = s
		i += incx

		x[i] = s
		i += incx

		x[i] = s
		i += incx
	}
	for n != 0 {
		x[i] = s
		i += incx
		n--
	}
	// k := n - 1
	// for ; k > 0; k -= 2 * incx {
	// 	x[k] = s
	// 	x[k-1] = s
	// }
	// if k == 0 {
	// 	x[0] = s
	// }
}
