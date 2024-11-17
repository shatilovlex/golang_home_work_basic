package shapes

import "fmt"

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) CalcSquare() float64 {
	return t.Base * t.Height / 2
}

func (t Triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %v, высота %v\n", t.Base, t.Height)
}
