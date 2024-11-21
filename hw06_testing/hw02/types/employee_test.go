package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployee_MakeCorrectString(t *testing.T) {
	expected := "User ID: 1; Age: 11; Name: Bil; Department ID: 22; "
	employee := Employee{
		UserID:       1,
		Age:          11,
		Name:         "Bil",
		DepartmentID: 22,
	}

	actual := employee.String()

	assert.Equal(t, expected, actual)
}
