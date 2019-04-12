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
			m.Col(j).Scale(v)
		}
	} else {
		for i := 0; i < m.nr; i++ {
			m.Row(i).Scale(v)
		}
	}
	return m
}
