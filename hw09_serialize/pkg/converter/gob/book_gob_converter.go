package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/shatilovlex/golang_home_work_basic/hw09_serialize/pkg/types"
)

func serialize(book *types.Book) ([]byte, error) {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	err := enc.Encode(book)
	return result.Bytes(), err
}

func unserialize(data []byte) (*types.Book, error) {
	var result types.Book
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	err := dec.Decode(&result)
	return &result, err
}
