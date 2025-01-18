package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	v := 0
	wg.Add(1)

	Worker(wg, mu, &v)

	assert.Equal(t, 1, v)
}
