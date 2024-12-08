package yaml

import (
	"github.com/shatilovlex/golang_home_work_basic/hw09_serialize/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serialize(t *testing.T) {
	tests := []struct {
		name string
		book types.Book
		want []byte
	}{
		{
			name: "test serialize successful",
			book: types.Book{
				Id:     1,
				Title:  "The Go Programming Language. - Addison-Wesley Professional",
				Size:   272,
				Rate:   9,
				Year:   2015,
				Author: "Donovan A., Kernighan В.",
				Sample: []byte("1"),
			},
			want: []byte("id: 1\nyear: 2015\nsize: 272\nrate: 9\n" +
				"title: The Go Programming Language. - Addison-Wesley Professional\n" +
				"author: Donovan A., Kernighan В.\nsample:\n    - 49\n"),
		},
		{
			name: "test serialize empty successful",
			book: types.Book{},
			want: []byte("{}\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serialize(tt.book)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_unserialize(t *testing.T) {
	book := []byte("id: 1\nyear: 2015\nsize: 272\nrate: 9\n" +
		"title: The Go Programming Language. - Addison-Wesley Professional\n" +
		"author: Donovan A., Kernighan В.\nsample:\n    - 49\n")
	want := &types.Book{
		Id:     1,
		Title:  "The Go Programming Language. - Addison-Wesley Professional",
		Size:   272,
		Rate:   9,
		Year:   2015,
		Author: "Donovan A., Kernighan В.",
		Sample: []byte("1"),
	}
	got, err := unserialize(book)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_unserialize_return_error_when_invalid_xml(t *testing.T) {
	book := []byte("-")
	_, err := unserialize(book)
	assert.Error(t, err)
}
