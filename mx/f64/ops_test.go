package f64

import (
	"testing"

	"bitbucket.org/vasthu/tnsr/blas"

	mx "github.com/c2akula/go.mx/mx/f64/blas"
)

func BenchmarkScale1(b *testing.B) {
	v := [100]float64{}
	for i := range v {
		v[i] = 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mx.Scale(len(v), v[:], 1, -1)
	}
}

func BenchmarkScale2(b *testing.B) {
	v := [100]float64{}
	for i := range v {
		v[i] = 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blas.Scale(len(v), v[:], 1, -1)
	}
}