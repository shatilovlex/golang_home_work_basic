package shapes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	types "hw06_testing/hw05/pkg/shape/types"
)

func TestShape_CalcSquare(t *testing.T) {
	tests := []struct {
		name  string
		shape types.Shape
		want  float64
	}{
		{
			name:  "Circle",
			shape: &Circle{Radius: 5},
			want:  78.53981633974483,
		},
		{
			name: "Rectangle",
			shape: &Rectangle{
				Width:  10,
				Height: 5,
			},
			want: float64(50),
		},
		{
			name: "Triangle",
			shape: &Triangle{
				Base:   8,
				Height: 6,
			},
			want: float64(24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.shape.CalcSquare())
		})
	}
}

func TestShape_String(t *testing.T) {
	tests := []struct {
		name  string
		shape fmt.Stringer
		want  string
	}{
		{
			name:  "Circle",
			shape: &Circle{Radius: 5},
			want:  "Круг: радиус 5\n",
		},
		{
			name: "Rectangle",
			shape: &Rectangle{
				Width:  10,
				Height: 5,
			},
			want: "Прямоугольник: ширина 10, высота 5\n",
		},
		{
			name: "Triangle",
			shape: &Triangle{
				Base:   8,
				Height: 6,
			},
			want: "Треугольник: основание 8, высота 6\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.shape.String())
		})
	}
}
