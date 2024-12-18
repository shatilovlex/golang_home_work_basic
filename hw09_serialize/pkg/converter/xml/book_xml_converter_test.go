package xml

import (
	"testing"

	"github.com/shatilovlex/golang_home_work_basic/hw09_serialize/pkg/types"
	"github.com/stretchr/testify/assert"
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
				ID:     1,
				Title:  "The Go Programming Language. - Addison-Wesley Professional",
				Size:   272,
				Rate:   9,
				Year:   2015,
				Author: "Donovan A., Kernighan В.",
				Sample: []byte("Sample"),
			},
			want: []byte("<Book><id>1</id><year>2015</year><size>272</size><rate>9</rate>" +
				"<title>The Go Programming Language. - Addison-Wesley Professional</title>" +
				"<author>Donovan A., Kernighan В.</author><sample>Sample</sample></Book>"),
		},
		{
			name: "test serialize empty successful",
			book: types.Book{},
			want: []byte("<Book></Book>"),
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
	book := []byte("<Book><id>1</id><year>2015</year><size>272</size><rate>9</rate>" +
		"<title>The Go Programming Language. - Addison-Wesley Professional</title>" +
		"<author>Donovan A., Kernighan В.</author><sample>Sample</sample></Book>")
	want := &types.Book{
		ID:     1,
		Title:  "The Go Programming Language. - Addison-Wesley Professional",
		Size:   272,
		Rate:   9,
		Year:   2015,
		Author: "Donovan A., Kernighan В.",
		Sample: []byte("Sample"),
	}
	got, err := unserialize(book)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_unserialize_return_error_when_invalid_xml(t *testing.T) {
	book := []byte("<Book</Book>")
	_, err := unserialize(book)
	assert.Error(t, err)
}
