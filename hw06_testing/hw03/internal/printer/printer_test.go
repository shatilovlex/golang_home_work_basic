package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateChessboard(t *testing.T) {
	tests := []struct {
		name string
		size int
		want string
	}{
		{
			name: "GetClassicChessboard",
			size: 8,
			want: " # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n",
		},
		{name: "GetSmallChessboard", size: 2, want: " #\n# \n"},
		{name: "GetOneSquare", size: 1, want: "#\n"},
		{name: "GetEmptyChessboard", size: 0, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateChessboard(tt.size)
			assert.Equal(t, tt.want, got)
		})
	}
}
