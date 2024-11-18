package shape

import (
	"errors"

	types "hw06_testing/hw05/pkg/shape/types"
)

var ErrObjectIsNotShape = errors.New("переданный объект не является фигурой")

func CalculateArea(s any) (float64, error) {
	if shape, ok := s.(types.Shape); ok {
		return shape.CalcSquare(), nil
	}

	return 0, ErrObjectIsNotShape
}
