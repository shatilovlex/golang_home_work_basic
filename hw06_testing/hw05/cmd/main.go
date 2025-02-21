package main

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw05/pkg/shape"
	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw05/pkg/shape/types"
	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw05/pkg/shape/types/shapes"
)

func main() {
	var circle, rectangle, triangle types.Shape
	circle = &shapes.Circle{Radius: 5}
	fmt.Print(shape.GetShapeString(circle))

	rectangle = &shapes.Rectangle{Width: 10, Height: 5}
	fmt.Print(shape.GetShapeString(rectangle))

	triangle = &shapes.Triangle{Base: 8, Height: 6}
	fmt.Print(shape.GetShapeString(triangle))

	noShape := "Я строка, а не фигура"
	fmt.Print(shape.GetShapeString(noShape))
}
