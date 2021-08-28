package goft

import (
	"math"
	"math/cmplx"
)

// DFT is a function to Discrete Fourier Transform.
func DFT(f interface{}) ([]complex128, error) {
	c, err := interfaceToComplexSlice(f)
	if err != nil {
		return nil, err
	}

	n := len(c)
	F := make([]complex128, n) // return value
	for t := 0; t < n; t++ {
		F[t] = dftMapping(c, float64(t))
	}
	return F, nil
}

func dftMapping(f []complex128, t float64) complex128 {
	var sumReal, sumImag float64
	N := float64(len(f))
	n := len(f)
	memoV := 2.0 * math.Pi / N * t
	for x := 0; x < n; x++ {
		v := memoV * float64(x)
		sinV, cosV := math.Sin(v), math.Cos(v)
		sumReal += real(f[x])*cosV + imag(f[x])*sinV
		sumImag += -real(f[x])*sinV + imag(f[x])*cosV
	}
	return complex(sumReal, sumImag)
}

// Amplitude is a function mainly to calc calculate the amplitude of the FFT result.
func Amplitude(c []complex128) []float64 {
	amp := make([]float64, len(c))
	for i := 0; i < len(c); i++ {
		amp[i] = cmplx.Abs(c[i])
	}
	return amp
}

// IDFT is a function to Inverse Discrete Fourier Transform.
func IDFT(F interface{}) ([]complex128, error) {
	c, err := interfaceToComplexSlice(F)
	if err != nil {
		return nil, err
	}

	n := len(c)
	f := make([]complex128, n) // return value
	for x := 0; x < n; x++ {
		f[x] = idftMapping(c, float64(x))
	}
	return f, nil
}

func idftMapping(F []complex128, x float64) complex128 {
	var sumReal, sumImag float64
	N := float64(len(F))
	n := len(F)
	memoV := 2.0 * math.Pi * x / N
	for t := 0; t < n; t++ {
		v := memoV * float64(t)
		sinV, cosV := math.Sin(v), math.Cos(v)
		sumReal += real(F[t])*cosV - imag(F[t])*sinV
		sumImag += real(F[t])*sinV + imag(F[t])*cosV
	}
	return complex(sumReal/N, sumImag/N)
}
