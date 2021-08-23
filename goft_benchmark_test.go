package goft

import (
	"testing"
)

func BenchmarkDft(b *testing.B) {
	tmp := []float64{0, 1, 0, -1}
	for i := 0; i < 10; i++ {
		tmp = append(tmp, tmp...)
	}
	DFT(tmp)
}
