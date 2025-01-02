package sensor

import (
	"math/rand"
	"time"
)

func Generator(sensorChan chan<- int, timer *time.Timer) {
	defer close(sensorChan)
wait:
	for {
		select {
		case <-timer.C:
			break wait
		default:
			time.Sleep(20 * time.Millisecond)
			sensorChan <- rand.Intn(100) + 1 //nolint:gosec
		}
	}
}
