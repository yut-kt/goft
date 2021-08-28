package goft

import (
	"errors"
)

type matrix [][]complex128

func multiMatrix(x, y matrix) (matrix, error) {
	xRowLen, xColumnLen := len(x), len(x[0])
	yRowLen, yColumnLen := len(y), len(y[0])
	if xColumnLen != yRowLen {
		return nil, errors.New("can not multi matrix size")
	}

	multiedMatrix := make(matrix, xRowLen)
	for i := 0; i < xRowLen; i++ {
		multiedMatrix[i] = make([]complex128, yColumnLen)
	}

	ch := make(chan int)

	for i := 0; i < xColumnLen; i++ {
		for j := 0; j < yRowLen; j++ {
			go multiEveryRowColumn(i, j, xColumnLen, x, y, multiedMatrix, ch)
		}
	}
	for i := 0; i < xColumnLen; i++ {
		for j := 0; j < yRowLen; j++ {
			<-ch
		}
	}
	return multiedMatrix, nil
}

func multiEveryRowColumn(
	i, j, calcCount int,
	x, y, returnMatrix matrix,
	ch chan int,
) {
	var sum complex128
	for k := 0; k < calcCount; k++ {
		sum += x[i][k] * y[k][j]
	}
	returnMatrix[i][j] = sum
	ch <- 0
}
