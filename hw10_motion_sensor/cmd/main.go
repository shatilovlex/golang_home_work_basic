package main

import (
	"fmt"
	"time"

	"github.com/shatilovlex/golang_home_work_basic/hw10_motion_sensor/internal/sensor"
)

func main() {
	results := make(chan sensor.MeanResult, 10)
	sensorChan := make(chan int, 10)
	timer := time.NewTimer(1 * time.Minute)

	go sensor.Generator(sensorChan, timer)
	go sensor.Worker(sensorChan, results)

	for result := range results {
		fmt.Println(result)
	}
}
