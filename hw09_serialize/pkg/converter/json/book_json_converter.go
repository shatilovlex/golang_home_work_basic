package json

import (
	"encoding/json"

	"github.com/shatilovlex/golang_home_work_basic/hw09_serialize/pkg/types"
)

func serialize(book types.Book) ([]byte, error) {
	return json.Marshal(book)
}

func unserialize(data []byte) (*types.Book, error) {
	var result types.Book
	err := json.Unmarshal(data, &result)
	return &result, err
}
