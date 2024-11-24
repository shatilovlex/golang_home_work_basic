package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countWords(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{name: "empty string", args: args{input: ""}, want: map[string]int{}},
		{
			name: "One word",
			args: args{input: "word"},
			want: map[string]int{
				"word": 1,
			},
		},
		{
			name: "simple text",
			args: args{input: "  \nLorem ipsum dolor sit amet, consectetur;"},
			want: map[string]int{
				"Lorem":       1,
				"ipsum":       1,
				"dolor":       1,
				"sit":         1,
				"amet":        1,
				"consectetur": 1,
			},
		},
		{
			name: "simple cyrillic text",
			args: args{input: "Высококачественный прототип будущего проекта, допускает внедрение новых принципов."},
			want: map[string]int{
				"Высококачественный": 1,
				"прототип":           1,
				"будущего":           1,
				"проекта":            1,
				"допускает":          1,
				"внедрение":          1,
				"новых":              1,
				"принципов":          1,
			},
		},
		{
			name: "simple correct calc count",
			args: args{input: "Test test test test"},
			want: map[string]int{
				"Test": 1,
				"test": 3,
			},
		},
		{
			name: "simple correct calc count",
			args: args{input: "Test test test test"},
			want: map[string]int{
				"Test": 1,
				"test": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.args.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
