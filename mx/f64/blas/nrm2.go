package blas

import "math"

func Nrm2(n int, x []float64, incx int) (nrm float64) {
	i := 0
	for ; n >= 4; n -= 4 {
		a := x[i]
		nrm += a * a
		i += incx

		b := x[i]
		nrm += b * b
		i += incx

		c := x[i]
		nrm += c * c
		i += incx

		d := x[i]
		nrm += d * d
		i += incx
	}

	for n != 0 {
		e := x[i]
		nrm += e * e
		i += incx
		n--
	}
	return math.Sqrt(nrm)
}
