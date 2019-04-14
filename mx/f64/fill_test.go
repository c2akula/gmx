package f64

import (
	"testing"

	"github.com/c2akula/go.mx/mx/f64/blas"
)

func (m *Mx) Fill1(s float64) *Mx {
	if m.IsVector() {
		if m.IsRow() {
			blas.Fill(m.nc, m.data, m.cs, s)
		} else {
			blas.Fill(m.nr, m.data, m.rs, s)
		}
		return m
	}
	if m.IsTranspose() {
		for j := 0; j < m.nc; j++ {
			// mc := m.Col(j)
			// blas.Fill(mc.len, mc.data, mc.rs, s)
			blas.Fill(m.nr, m.col(j), m.rs, s)
		}
	} else {
		for i := 0; i < m.nr; i++ {
			// mr := m.Row(i)
			// blas.Fill(m.nc, mr.data, mr.cs, s)
			blas.Fill(m.nc, m.row(i), m.cs, s)
		}
	}
	return m
}

func BenchmarkMx_Fill1(b *testing.B) {
	m := randmx(100, 100) // .Transpose()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Fill1(-1)
	}
}

func TestFill1_NoBlk_NT_T(t *testing.T) {
	m := New([]float64{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
		16, 17, 18, 19,
	}, 5, 4)
	tests := []struct {
		got, exp *Mx
	}{
		{
			got: m.Dup().Fill1(-1),
			exp: Ones(m.nr, m.nc).Scale(-1),
		},
		{
			got: m.Dup().Transpose().Fill1(-1),
			exp: New([]float64{
				-1, -1, -1, -1, -1,
				-1, -1, -1, -1, -1,
				-1, -1, -1, -1, -1,
				-1, -1, -1, -1, -1,
				-1, -1, -1, -1, -1,
			}, 4, 5),
		},
	}
	for ti, tt := range tests {
		if !IsEqual(tt.got, tt.exp) {
			t.Fatalf("test[%d]: failed! got=%s, exp=%s", ti, tt.got, tt.exp)
		}
	}
}

func TestFill1_Blk_NT(t *testing.T) {
	m := New([]float64{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
		16, 17, 18, 19,
	}, 5, 4)
	m.Blk(2, 1, 2, 2).Fill1(-1)
	tests := []struct {
		got, exp *Mx
	}{
		{
			got: m,
			exp: New([]float64{
				0, 1, 2, 3,
				4, 5, 6, 7,
				8, -1, -1, 11,
				12, -1, -1, 15,
				16, 17, 18, 19,
			}, 5, 4),
		},
	}
	for ti, tt := range tests {
		if !IsEqual(tt.got, tt.exp) {
			t.Fatalf("test[%d]: failed! got=%s, exp=%s", ti, tt.got, tt.exp)
		}
	}
}

func (m *Mx) Fill3(s float64) *Mx {
	if m.IsVector() {
		if m.IsRow() {
			j := m.nr - 1
			for ; j > 0; j -= 2 {
				m.data[j] = s
				m.data[j-1] = s
			}
			if j == 0 {
				m.data[0] = s
			}
		} else {
			i := m.nc - 1
			for ; i > 0; i -= 2 {
				m.data[i] = s
				m.data[i-1] = s
			}

			if i == 0 {
				m.data[0] = s
			}
		}
		return m
	}

	if !m.IsTranspose() {
		for i := 0; i < m.nr; i++ {
			mr := m.row(i)
			k := m.nc - 1
			for ; k > 0; k -= 2 {
				mr[k] = s
				mr[k-1] = s
			}

			if k == 0 {
				mr[0] = s
			}
		}
	} else {
		for j := 0; j < m.nc; j++ {
			mc := m.col(j)
			k := m.nr - 1
			for ; k > 0; k -= 2 {
				mc[k] = s
				mc[k-1] = s
			}
			if k == 0 {
				mc[0] = s
			}
		}
	}

	return m
}

func TestFill3_NoBlk_NT_T(t *testing.T) {
	m := New([]float64{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
		16, 17, 18, 19,
	}, 5, 4)
	tests := []struct {
		got, exp *Mx
	}{
		{
			got: m.Dup().Fill3(-1),
			exp: Ones(m.nr, m.nc).Scale(-1),
		},
		{
			got: m.Dup().Transpose().Fill3(-1),
			exp: Ones(5, 4).Transpose().Scale(-1),
		},
	}
	for ti, tt := range tests {
		if !IsEqual(tt.got, tt.exp) {
			t.Fatalf("test[%d]: failed! got=%s, exp=%s", ti, tt.got, tt.exp)
		}
	}
}

func TestFill3_Blk_NT(t *testing.T) {
	m := New([]float64{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
		16, 17, 18, 19,
	}, 5, 4)
	m.Blk(2, 1, 2, 2).Fill3(-1)
	got := m
	exp := New([]float64{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, -1, -1, 11,
		12, -1, -1, 15,
		16, 17, 18, 19,
	}, 5, 4)
	if !IsEqual(got, exp) {
		t.Fatalf("test failed! got=%s, exp=%s", got, exp)
	}
}

func BenchmarkFill3(b *testing.B) {
	m := randmx(100, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m = m.Fill3(-1)
	}
}
