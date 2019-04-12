package f64

type Submx struct {
	Mx
}

type Rng struct {
	b, e int
}

func (m *Mx) Slc(i, j Rng) *Submx {
	s := &Submx{}
	s.data = m.data[m.Sub2ind(i.b, j.b) : m.Sub2ind(i.e, j.e)+1]
	s.nr, s.nc = i.e-i.b+1, j.e-j.b+1
	s.rs, s.cs = m.rs, m.cs
	s.len = s.nr * s.nc
	return s
}

// Row extracts the ith row from matrix m
func (m *Mx) Row(i int) *Submx { return m.Slc(Rng{b: i, e: i}, Rng{b: 0, e: m.nc - 1}) }

// Col extracts the jth row from matrix m
func (m *Mx) Col(j int) *Submx { return m.Slc(Rng{b: 0, e: m.nr - 1}, Rng{b: j, e: j}) }

// Blk extracts the submatrix of size (nr,nc) starting at m(i,j)
func (m *Mx) Blk(i, j int, nr, nc int) *Submx {
	return m.Slc(Rng{b: i, e: i + nr - 1}, Rng{b: j, e: j + nc - 1})
}

func (m *Mx) Diag(k int) *Submx {
	n := mini(m.nc, m.nc)
	di, dj := 0, 0
	if k < 0 {
		di = -k
		n = mini(m.nr+k, m.nc)
	} else {
		dj = k
		n = mini(m.nr, m.nc-k)
	}
	s := &Submx{}
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
