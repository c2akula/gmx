package main

import (
	"fmt"

	"github.com/c2akula/go.mx/mx/f64"
)

func main() {
	m := f64.Eye(5, 4).Diag(0).Scale(-4)
	fmt.Println("m: ", m)

}
