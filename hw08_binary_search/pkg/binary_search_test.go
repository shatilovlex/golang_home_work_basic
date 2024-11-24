package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "not found value", args: args{arr: []int{3, 4}, item: 1}, want: -1},
		{name: "not found value in unsorted array", args: args{arr: []int{3, 2, 1}, item: 1}, want: -1},
		{name: "found value in sorted array", args: args{arr: []int{1, 2, 3}, item: 2}, want: 1},
		{name: "found first value in sorted array", args: args{arr: []int{1, 2, 3}, item: 1}, want: 0},
		{name: "found last value in sorted array", args: args{arr: []int{1, 2, 3}, item: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.args.arr, tt.args.item)
			assert.Equal(t, tt.want, got)
		})
	}
}
