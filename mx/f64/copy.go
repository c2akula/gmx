package f64

import "fmt"

func (m *Mx) Dup() *Mx {
	s := Zeros(m.nr, m.nc)
	copy(s.data, m.data)
	return s
}

func Copy(dst, src *Mx) {
	if dst.nr != src.nr || dst.nc != src.nc {
		panic(fmt.Errorf("mx:Copy: dimensions mismatch"))
	}

	st, dt := src.IsTranspose(), dst.IsTranspose()
	switch {
	case st && dt:
		for j := 0; j < src.nc; j++ {
			copy(dst.col(j), src.col(j))
		}
	case st && !dt:
		for j := 0; j < src.nc; j++ {
			for i, sij := range src.col(j) {
				dst.Setij(i, j, sij)
			}
		}
	case !st && dt:
		for i := 0; i < src.nr; i++ {
			for j, sij := range src.row(i) {
				dst.Setij(i, j, sij)
			}
		}
	default:
		for i := 0; i < src.nr; i++ {
			copy(dst.row(i), src.row(i))
		}
	}
}
