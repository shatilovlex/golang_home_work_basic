package sensor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker_CanProcessTenValues(t *testing.T) {
	inputChan := make(chan int, 10)
	outputChan := make(chan MeanResult, 1)
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	meanResult := 5.5
	for _, v := range input {
		inputChan <- v
	}
	close(inputChan)

	Worker(inputChan, outputChan)

	result, ok := <-outputChan
	assert.True(t, ok)
	assert.Equal(t, meanResult, result.Result)
}

func TestWorker_NoResultForLessTenValues(t *testing.T) {
	inputChan := make(chan int, 10)
	outputChan := make(chan MeanResult, 1)
	input := []int{1, 2, 3, 4, 5}
	for _, v := range input {
		inputChan <- v
	}
	close(inputChan)

	Worker(inputChan, outputChan)

	_, ok := <-outputChan
	assert.False(t, ok)
}

func TestWorker_NoResultForEmptyChanel(t *testing.T) {
	inputChan := make(chan int, 10)
	outputChan := make(chan MeanResult, 1)
	close(inputChan)

	Worker(inputChan, outputChan)

	_, ok := <-outputChan
	assert.False(t, ok)
}
