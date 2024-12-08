package message_pack

import (
	"github.com/shatilovlex/golang_home_work_basic/hw09_serialize/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serialize_unserialize_book(t *testing.T) {
	book := &types.Book{
		Id:     1,
		Title:  "The Go Programming Language. - Addison-Wesley Professional",
		Size:   272,
		Rate:   9,
		Year:   2015,
		Author: "Donovan A., Kernighan В.",
		Sample: []byte("Sample"),
	}
	serializedBook, err := serialize(book)
	assert.NoError(t, err)
	got, err := unserialize(serializedBook)
	assert.NoError(t, err)
	assert.Equal(t, book, got)
}

func Test_unserialize_return_error_when_invalid_xml(t *testing.T) {
	book := []byte{0x01}
	_, err := unserialize(book)
	assert.Error(t, err)
}
