package f64

import "github.com/c2akula/go.mx/mx/f64/blas"

func (m *Mx) Scale(v float64) *Mx {
	if m.IsVector() {
		if m.IsRow() {
			blas.Scale(len(m.data), m.data, m.cs, v)
		} else {
			blas.Scale(len(m.data), m.data, m.rs, v)
		}
		return m
	}

	if m.IsTranspose() {
		for j := 0; j < m.nc; j++ {
			mc := m.Col(j)
			for i := 0; i < m.nr; i++ {
				mc.data[mc.Sub2ind(i, 0)] *= v
			}
		}
	} else {
		for i := 0; i < m.nr; i++ {
			mr := m.Row(i)
			for j := 0; j < m.nc; j++ {
				mr.data[mr.Sub2ind(0, j)] *= v
			}
		}
	}
	return m
}
