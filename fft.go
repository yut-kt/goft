package goft

import (
	"errors"
	"math"
	"math/cmplx"
)

type Type int

const (
	CaseFFT Type = iota
	CaseIFFT
)

// FFT is a function to Fast Fourier Transformation.
func FFT(f interface{}) ([]complex128, error) {
	c, err := interfaceToComplexSlice(f)
	if err != nil {
		return nil, err
	}

	if l := len(c); (l & (l - 1)) != 0 {
		return nil, errors.New("must len(input) == 2^x")
	}

	recursive(c, CaseFFT)
	return c, err
}

// IFFT is a function to Inverse Fast Fourier Transform.
func IFFT(F interface{}) ([]complex128, error) {
	c, err := interfaceToComplexSlice(F)
	if err != nil {
		return nil, err
	}

	if l := len(c); (l & (l - 1)) != 0 {
		return nil, errors.New("must len(input) == 2^x")
	}

	recursive(c, CaseIFFT)
	N := float64(len(c))
	for i, v := range c {
		c[i] = complex(real(v)/N, imag(v)/N)
	}
	return c, err
}

func recursive(c []complex128, t Type) {
	n := len(c)
	if n == 1 {
		return
	}
	halfN := n / 2
	even := make([]complex128, halfN)
	for i := 0; i < n; i += 2 {
		even[i/2] = c[i]
	}
	odd := make([]complex128, halfN)
	for i := 1; i < n; i += 2 {
		odd[i/2] = c[i]
	}
	switch t {
	case CaseFFT:
		recursive(even, CaseFFT)
		recursive(odd, CaseFFT)
		calcFFT(c, even, odd)
	case CaseIFFT:
		recursive(even, CaseIFFT)
		recursive(odd, CaseIFFT)
		calcIFFT(c, even, odd)
	}
}

func calcFFT(c, even, odd []complex128) {
	n := len(c)
	halfN := n / 2
	v := -2.0 * math.Pi
	for i := 0; i < n; i++ {
		c[i] = even[i%halfN] + odd[i%halfN]*cmplx.Exp(complex(0, v*float64(i)/float64(n)))
	}
}

func calcIFFT(c, even, odd []complex128) {
	n := len(c)
	halfN := n / 2
	v := 2.0 * math.Pi
	for i := 0; i < n; i++ {
		c[i] = even[i%halfN] + odd[i%halfN]*cmplx.Exp(complex(0, v*float64(i)/float64(n)))
	}
}
