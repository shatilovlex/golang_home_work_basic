package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw02/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	return data, err
}
