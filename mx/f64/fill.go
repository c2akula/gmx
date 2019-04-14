package f64

func (m *Mx) Fill(s float64) *Mx {
	for i := 0; i < m.nr; i++ {
		mr := m.Row(i)
		for j := 0; j < m.nc; j++ {
			mr.Set(j, s)
		}
	}
	return m
}
