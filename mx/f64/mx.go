package f64

import (
	"fmt"
)

type Mx struct {
	data   []float64
	rs, cs int
	nr, nc int
	len    int
}

func New(data []float64, nr, nc int) *Mx {
	m := &Mx{
		data: make([]float64, nr*nc),
		nr:   nr,
		nc:   nc,
		rs:   nc,
		cs:   1,
		len:  nr * nc,
	}
	if data != nil {
		copy(m.data, data)
	}
	return m
}

func Zeros(nr, nc int) *Mx {
	return &Mx{
		data: make([]float64, nr*nc),
		nr:   nr,
		nc:   nc,
		rs:   nc,
		cs:   1,
		len:  nr * nc,
	}
}

func Ones(nr, nc int) *Mx {
	m := Zeros(nr, nc)
	for i := range m.data {
		m.data[i] = 1
	}
	return m
}

func Eye(nr, nc int) *Mx {
	m := Zeros(nr, nc)
	for i := 0; i < m.len; i += m.rs + m.cs {
		m.data[i] = 1
	}
	return m
}

func Diag(v []float64, k int) *Mx {
	n := len(v)
	if k > 0 {
		n += k
	} else {
		n -= k
	}
	m := Zeros(n, n)
	if k >= 0 {
		for i, e := range v {
			m.data[m.Sub2ind(i, i)+k] = e
		}
	} else {
		for i, e := range v {
			m.data[m.Sub2ind(i-k, i-k)+k] = e
		}
	}
	return m
}

func (m *Mx) IsVector() bool    { return m.nr == 1 || m.nc == 1 }
func (m *Mx) IsRow() bool       { return m.nr == 1 }
func (m *Mx) IsCol() bool       { return m.nc == 1 }
func (m *Mx) IsTranspose() bool { return m.rs == 1 }

func (m *Mx) Transpose() *Mx {
	m.nr, m.nc = m.nc, m.nr
	m.rs, m.cs = m.cs, m.rs
	return m
}

func (m *Mx) Size() (nr, nc int) { return m.nr, m.nc }
func (m *Mx) Len() (n int)       { return m.len }

func (m *Mx) String() string {
	s := ""
	for i := 0; i < m.nr; i++ {
		s += "["
		for j := 0; j < m.nc; j++ {
			// fmt.Printf("[%d,%d]: %d\n", i, j, m.Sub2ind(i, j))
			s += fmt.Sprintf(" %8.4f ", m.Getij(i, j))
		}
		s += "]\n"
	}
	return s
}
