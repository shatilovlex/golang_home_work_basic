package yaml

import (
	"github.com/shatilovlex/golang_home_work_basic/hw09_serialize/pkg/types"
	"gopkg.in/yaml.v3"
)

func serialize(book types.Book) ([]byte, error) {
	return yaml.Marshal(book)
}

func unserialize(book []byte) (*types.Book, error) {
	var result types.Book
	var err error
	err = yaml.Unmarshal(book, &result)

	return &result, err
}
