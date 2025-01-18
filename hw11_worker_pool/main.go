package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	const grCount = 10000
	v := 0
	var mu sync.Mutex
	for i := 0; i < grCount; i++ {
		wg.Add(1)
		go Worker(&wg, &mu, &v)
	}
	wg.Wait()
	fmt.Print(v)
}

func Worker(wg *sync.WaitGroup, mu *sync.Mutex, v *int) {
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	*v++
}
