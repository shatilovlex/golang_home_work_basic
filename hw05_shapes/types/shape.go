package types

import (
	"math"
)

type Shape interface {
	CalculateArea() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) CalculateArea() float64 {
	return math.Pi * c.Radius / 2
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) CalculateArea() float64 {
	return t.Base * t.Height / 2
}
