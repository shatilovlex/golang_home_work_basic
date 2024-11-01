package shape

import (
	"errors"

	"github.com/shatilovlex/golang_home_work_basic/hw05_shapes/pkg/shape/types"
)

func CalculateArea(s any) (float64, error) {
	if shape, ok := s.(types.Shape); ok {
		return shape.CalcSquare(), nil
	}

	return 0, errors.New("переданный объект не является фигурой")
}
