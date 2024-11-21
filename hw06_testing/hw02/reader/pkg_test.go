package reader

import (
	"testing"

	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw02/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadJSON(t *testing.T) {
	result, err := ReadJSON("validJson.json")

	require.NoError(t, err)
	subset := types.Employee{Name: "Rob", UserID: 10, Age: 25, DepartmentID: 3}
	assert.Contains(t, result, subset)
}

func TestReadJSONForInvalidJson(t *testing.T) {
	_, err := ReadJSON("invalidJson.json")

	assert.ErrorContains(t, err, "unexpected end of JSON input")
}

func TestReadJSONForNotExistsFile(t *testing.T) {
	_, err := ReadJSON("notExistsFile.json")

	assert.ErrorContains(t, err, "open notExistsFile.json: no such file or directory")
}
