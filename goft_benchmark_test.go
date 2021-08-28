package goft

import (
	"math"
	"testing"
)

func BenchmarkDft(b *testing.B) {
	N := 1024
	v := make([]float64, N)
	circle := 2.0 * math.Pi
	for i := 0; i < N; i++ {
		v[i] = math.Sin(float64(i) * circle)
	}
	DFT(v)
}

func BenchmarkFFT(b *testing.B) {
	N := 1024
	v := make([]float64, N)
	circle := 2.0 * math.Pi
	for i := 0; i < N; i++ {
		v[i] = math.Sin(float64(i) * circle)
	}
	FFT(v)
}
