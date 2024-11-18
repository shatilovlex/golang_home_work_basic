package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	shapes "hw06_testing/hw05/pkg/shape/types/shapes"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want float64
	}{
		{
			name: "circle",
			arg:  &shapes.Circle{Radius: 5},
			want: 78.53981633974483,
		},
		{
			name: "Rectangle",
			arg: &shapes.Rectangle{
				Width:  10,
				Height: 5,
			},
			want: float64(50),
		},
		{
			name: "Triangle",
			arg: &shapes.Triangle{
				Base:   8,
				Height: 6,
			},
			want: float64(24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateArea(tt.arg)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculateAreaError(t *testing.T) {
	_, err := CalculateArea("")
	require.ErrorIs(t, err, ErrObjectIsNotShape)
}
