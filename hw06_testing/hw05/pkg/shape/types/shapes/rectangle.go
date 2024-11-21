package shapes

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) CalcSquare() float64 {
	return r.Width * r.Height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %v, высота %v\n", r.Width, r.Height)
}
