package f64

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSubmx(t *testing.T) {
	m := Zeros(5, 4)
	for i := 0; i < m.len; i++ {
		m.data[i] = float64(rand.Intn(m.len))
	}
	fmt.Println(m)

	// get submatrix
	s := m.Slc(Rng{1, 3}, Rng{1, 2})
	fmt.Println("s:", s)
	// set submatrix
	s.Setij(1, 1, -1)
	fmt.Println("s':", s)
	fmt.Println("m':", m)

	ss := s.Slc(Rng{1, 2}, Rng{0, 1})
	fmt.Println(ss)
	ss.Setij(1, 1, -1)
	fmt.Println("ss':", ss)
	fmt.Println("m':", m)

	sss := ss.Slc(Rng{0, 1}, Rng{0, 0})
	fmt.Println(sss)
	sss.Setij(1, 0, -1)
	fmt.Println("sss':", sss)
	fmt.Println("m':", m)
}

func TestMx_Blk(t *testing.T) {
	m := Zeros(5, 4)
	for i := 0; i < m.len; i++ {
		m.data[i] = float64(rand.Intn(m.len))
	}
	fmt.Println(m)

	s := m.Blk(1,1, 3, 2)
	fmt.Println("s: ", s)
}

func TestMx_Diag(t *testing.T) {
	m := Zeros(5, 4)
	for i := 0; i < m.len; i++ {
		m.data[i] = float64(rand.Intn(m.len))
	}
	fmt.Println(m)

	s := m.Diag(2)
	fmt.Println("s: ", s)
}