package shape

import (
	"testing"

	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw05/pkg/shape/types/shapes"
	"github.com/stretchr/testify/assert"
)

func TestGetShapeString(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want string
	}{
		{
			name: "circle",
			arg:  &shapes.Circle{Radius: 5},
			want: "Круг: радиус 5\nПлощадь: 78.53981633974483",
		},
		{
			name: "Rectangle",
			arg: &shapes.Rectangle{
				Width:  10,
				Height: 5,
			},
			want: "Прямоугольник: ширина 10, высота 5\nПлощадь: 50",
		},
		{
			name: "Triangle",
			arg: &shapes.Triangle{
				Base:   8,
				Height: 6,
			},
			want: "Треугольник: основание 8, высота 6\nПлощадь: 24",
		},
		{
			name: "any",
			arg:  "any",
			want: "Ошибка: переданный объект не является фигурой.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetShapeString(tt.arg))
		})
	}
}
