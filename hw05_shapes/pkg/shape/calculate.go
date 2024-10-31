package shape

import (
	"errors"
	"github.com/shatilovlex/golang_home_work_basic/hw05_shapes/pkg/shape/types/shapes"
)

func CalculateArea(s any) (float64, error) {
	var area float64
	var err error
	switch shape := s.(type) {
	case shapes.Circle:
		area = shape.CalcSquare()
	case shapes.Rectangle:
		area = shape.CalcSquare()
	case shapes.Triangle:
		area = shape.CalcSquare()
	default:
		err = errors.New("переданный объект не является фигурой")
	}
	return area, err
}
