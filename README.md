# goft

[![v0.1.0](https://img.shields.io/github/v/release/yut-kt/goft?logoColor=ff69b4&style=social)]()
[![Test](https://github.com/yut-kt/goft/actions/workflows/default_branch_test.yaml/badge.svg)](https://github.com/yut-kt/gowave/actions/workflows/default_branch_test.yaml)
[![coverage](https://img.shields.io/badge/coverage-75.2%25-green)](https://raw.githubusercontent.com/yut-kt/gowave/main/coverage/v0.1.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/yut-kt/goft)](https://goreportcard.com/report/github.com/yut-kt/gowave)  
[![Go Reference](https://pkg.go.dev/badge/github.com/yut-kt/goft.svg)](https://pkg.go.dev/github.com/yut-kt/goft)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/yut-kt/goft/main/LICENSE)


**Fourier Transformation support for Go language**

## Install
```bash
$ go get github.com/yut-kt/goft
```

## Usage
```go
import (
    "fmt"
    "math"
	
    "github.com/yut-kt/goft"
)

func ExampleDFT() {
    testSamples := []float64{0, 1, 0, -1}

    // DFT Any Numeric Slice
    f, err := goft.DFT(testSamples)
    if err != nil {
        panic(err)
    }
}
```

See [gowave_examples_test.go](https://github.com/yut-kt/goft/blob/main/goft_test.go) for detailed Usage

## License
goft is released under the [MIT License](https://raw.githubusercontent.com/yut-kt/goft/main/LICENSE).
