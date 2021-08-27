package goft_test

import (
	"fmt"
	"math"

	"github.com/yut-kt/goft"
)

const EPSILON float64 = 0.00000001

func ExampleDFT() {
	testSamples := []float64{0, 1, 0, -1}

	// DFT
	f, err := goft.DFT(testSamples)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	// Output:
	// [(0+0i) (2.449293598294703e-16-2i) (0+2.449293598294703e-16i) (-7.347880794884109e-16+2i)]
}

func ExampleAmplitude() {
	testSamples := []float64{0, 1, 0, -1}

	// DFT
	f, err := goft.DFT(testSamples)
	if err != nil {
		panic(err)
	}

	fmt.Println(goft.Amplitude(f))
	// Output:
	// [0 2 2.449293598294703e-16 2]
}

func ExampleIDFT() {
	testSamples := []float64{0, 1, 0, -1}

	// DFT
	f, err := goft.DFT(testSamples)
	if err != nil {
		panic(err)
	}

	// IDFT
	F, err := goft.IDFT(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(F)

	// float64 approximate check
	isEqual := true
	for i, v := range F {
		isEqual = isEqual && math.Abs(testSamples[i]-real(v)) < EPSILON
	}
	fmt.Println(isEqual)
	// Output:
	// [(-1.2246467991473515e-16+5.551115123125783e-17i) (1+6.123233995736757e-17i) (1.4997597826618535e-32+1.1102230246251565e-16i) (-1+6.123233995736757e-17i)]
	// true
}
