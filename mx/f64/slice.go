package f64

// Row extracts the ith row from matrix m
func (m *Mx) Row(i int) *Mx { return m.Blk(i, 0, 1, m.nc) }

// Col extracts the jth row from matrix m
func (m *Mx) Col(j int) *Mx { return m.Blk(0, j, m.nr, 1) }

// Blk extracts the submatrix of size (nr,nc) starting at m(i,j)
func (m *Mx) Blk(i, j int, nr, nc int) *Mx {
	s := &Mx{}
	s.data = m.data[m.Sub2ind(i, j):]
	s.nr, s.nc = nr, nc
	s.rs, s.cs = m.rs, m.cs
	s.len = s.nr * s.nc
	return s
}

func (m *Mx) Diag(k int) *Mx {
	n := mini(m.nc, m.nc)
	di, dj := 0, 0
	if k < 0 {
		di = -k
		n = mini(m.nr+k, m.nc)
	} else {
		dj = k
		n = mini(m.nr, m.nc-k)
	}
	s := &Mx{}
	s.data = m.data[m.Sub2ind(di, dj):]
	s.rs, s.cs = n, m.cs+m.nc
	s.nr, s.nc = 1, n
	s.len = s.nr * s.nc
	return s
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m *Mx) row(i int) []float64 {
	rb, re := m.rbe(i)
	return m.data[rb:re]
}

func (m *Mx) col(j int) []float64 {
	cb, ce := m.cbe(j)
	return m.data[cb:ce]
}
