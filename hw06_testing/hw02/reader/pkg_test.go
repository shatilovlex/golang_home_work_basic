package reader

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"

	"hw06_testing/hw02/types"
)

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestReadJSON(t *testing.T) {
	result, err := ReadJSON("data.json")

	require.NoError(t, err)
	subset := types.Employee{Name: "Rob", UserID: 10, Age: 25, DepartmentID: 3}
	assert.Contains(t, result, subset)
}
func TestReadJSONError(t *testing.T) {
	_, err := ReadJSON("emptyFile.json")

	assert.ErrorIs(t, err, JsonInvalidError)
}

func TestGetFile(t *testing.T) {
	file, err := getFile("emptyFile.json")

	require.NoError(t, err)
	assert.NotNil(t, file)
}

func TestGetFileError(t *testing.T) {
	_, err := getFile("")

	require.ErrorIs(t, err, NoSuchFileOrDirectoryError)
}

func TestReadBytes(t *testing.T) {
	input := strings.NewReader("Json data")

	bytes, err := readBytes(input)

	require.NoError(t, err)
	assert.NotNil(t, bytes)
}

func TestReadBytesError(t *testing.T) {
	input := errReader{}

	_, err := readBytes(input)

	require.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadEmployeesByBytes(t *testing.T) {
	input := `[{"userId": 10,"age": 25,"name": "Rob","departmentId": 3}]`

	employees, err := getEmployeesByBytes([]byte(input))

	require.NoError(t, err)
	subset := types.Employee{Name: "Rob", UserID: 10, Age: 25, DepartmentID: 3}
	assert.Contains(t, employees, subset)

}
func TestReadEmployeesByBytesError(t *testing.T) {
	input := `]`

	_, err := getEmployeesByBytes([]byte(input))

	require.Error(t, err)
}
