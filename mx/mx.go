package main

import (
	"fmt"
	"math/rand"

	"github.com/c2akula/go.mx/mx/f64"
)

func main() {
	m := f64.Zeros(5, 4)
	for i := 0; i < m.Len(); i++ {
		m.Set(i, float64(rand.Intn(m.Len())))
	}
	fmt.Println(m)
	n := m.Scale(0.5)
	fmt.Println(n)
}
