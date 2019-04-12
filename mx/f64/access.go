package f64

type Indexer interface {
	Sub2ind(i, j int) (k int)
	Ind2sub(k int) (i, j int)
}

type GetSetter interface {
	Get(i int) float64
	Getij(i, j int) float64
	Set(i int, v float64)
	Setij(i, j int, v float64)
}

func (m *Mx) Sub2ind(i, j int) (k int) { k = i*m.rs + j*m.cs; return }
func (m *Mx) Ind2sub(k int) (i, j int) { i, j = k/m.nc, k%m.nc; return }

func (m *Mx) Get(i int) float64         { return m.data[m.Sub2ind(m.Ind2sub(i))] }
func (m *Mx) Set(i int, v float64)      { m.data[m.Sub2ind(m.Ind2sub(i))] = v }
func (m *Mx) Getij(i, j int) float64    { return m.data[m.Sub2ind(i, j)] }
func (m *Mx) Setij(i, j int, v float64) { m.data[m.Sub2ind(i, j)] = v }

// rbe returns the (b)egin and (e)nd of row i
func (m *Mx) rbe(i int) (b, e int) {
	b = i * m.rs
	e = b + m.cs*(m.nc-1) + 1
	return
}
// cbe returns the (b)egin and (e)nd of col j
func (m *Mx) cbe(j int) (b, e int) {
	b = j * m.cs
	e = b + m.rs*(m.nr-1) + 1
	return
}

// mbe returns the (b)egin and (e)nd of matrix m
func (m *Mx) mbe() (b, e int) {
	e = m.rs*(m.nr-1) + m.cs*(m.nc-1)+1
	return
}