package goft

import (
	"errors"
	"reflect"
)

func interfaceToComplexSlice(i interface{}) ([]complex128, error) {
	if reflect.ValueOf(i).Kind() != reflect.Slice {
		return nil, errors.New("interface is not slice")
	}

	var cSlice []complex128
	switch slice := i.(type) {
	case []complex128:
		cSlice = append(cSlice, slice...)
	case []float64:
		for _, v := range slice {
			cSlice = append(cSlice, complex(v, 0))
		}
	default:
		return nil, errors.New("unsupported type")
	}
	return cSlice, nil
}
