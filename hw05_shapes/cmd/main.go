package main

import (
	"github.com/shatilovlex/golang_home_work_basic/hw05_shapes/pkg/shape"
	"github.com/shatilovlex/golang_home_work_basic/hw05_shapes/pkg/shape/types"
	"github.com/shatilovlex/golang_home_work_basic/hw05_shapes/pkg/shape/types/shapes"
)

func main() {
	var circle, rectangle, triangle types.Shape
	circle = &shapes.Circle{Radius: 5}
	shape.PrintShape(circle)

	rectangle = &shapes.Rectangle{Width: 10, Height: 5}
	shape.PrintShape(rectangle)

	triangle = &shapes.Triangle{Base: 8, Height: 6}
	shape.PrintShape(triangle)

	noShape := "Я строка, а не фигура"
	shape.PrintShape(noShape)
}
