package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) CalcSquare() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Круг: радиус %v\n", c.Radius)
}
