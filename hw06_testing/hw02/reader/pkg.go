package reader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"hw06_testing/hw02/types"
)

var NoSuchFileOrDirectoryError = errors.New("no such file or directory")
var JsonInvalidError = errors.New("json invalid error")

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := getFile(
		filePath)
	if err != nil {
		return nil, fmt.Errorf("get JSON failed: %w", err)
	}

	bytes, err := readBytes(f)
	if err != nil {
		return nil, err
	}
	return getEmployeesByBytes(bytes)
}

func getEmployeesByBytes(bytes []byte) ([]types.Employee, error) {
	var data []types.Employee
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, JsonInvalidError
	}
	return data, nil
}

func readBytes(r io.Reader) ([]byte, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getFile(filePath string) (io.Reader, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, NoSuchFileOrDirectoryError
	}
	return f, nil
}
