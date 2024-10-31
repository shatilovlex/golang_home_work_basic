package main

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw05_shapes/types"
)

func calculateArea(s any) {
	switch s.(type) {
	case types.Circle:
		circle := s.(types.Circle)
		fmt.Printf("Круг: радиус %v\nПлощадь: %v\n\n", circle.Radius, circle.CalculateArea())
	case types.Rectangle:
		rectangle := s.(types.Rectangle)
		fmt.Printf(
			"Прямоугольник: ширина %v, высота %v\nПлощадь: %v\n\n",
			rectangle.Width,
			rectangle.Height,
			rectangle.CalculateArea(),
		)
	case types.Triangle:
		triangle := s.(types.Triangle)
		fmt.Printf(
			"Треугольник: основание %v, высота %v\nПлощадь: %v\n\n",
			triangle.Base,
			triangle.Height,
			triangle.CalculateArea(),
		)
	default:
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	}
}

func main() {
	circle := types.Circle{Radius: 5}
	calculateArea(circle)

	rectangle := types.Rectangle{Width: 10, Height: 5}
	calculateArea(rectangle)

	triangle := types.Triangle{Base: 8, Height: 6}
	calculateArea(triangle)

	calculateArea("Я строка, а не фигура")
}
