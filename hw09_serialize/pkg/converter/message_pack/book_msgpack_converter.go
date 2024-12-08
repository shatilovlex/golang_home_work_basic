package message_pack

import (
	"github.com/shatilovlex/golang_home_work_basic/hw09_serialize/pkg/types"
	"github.com/vmihailenco/msgpack/v5"
)

func serialize(book *types.Book) ([]byte, error) {
	return msgpack.Marshal(book)
}

func unserialize(data []byte) (*types.Book, error) {
	var result types.Book
	err := msgpack.Unmarshal(data, &result)
	return &result, err
}
