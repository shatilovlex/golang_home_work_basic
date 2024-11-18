package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleReader(t *testing.T) {
	result := ConsoleReader()

	assert.Equal(t, 8, result)
}
