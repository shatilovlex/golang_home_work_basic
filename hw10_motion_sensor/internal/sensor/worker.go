package sensor

import "fmt"

type MeanResult struct {
	SensorValues []float64
	Result       float64
}

func (m MeanResult) String() string {
	return fmt.Sprintf("MeanResult(%.0f): %.1f", m.SensorValues, m.Result)
}

func Worker(inputChan <-chan int, outputChan chan<- MeanResult) {
	defer close(outputChan)
	sensorValues := make([]float64, 0, 10)
	for v := range inputChan {
		sensorValues = append(sensorValues, float64(v))
		if len(sensorValues) == 10 {
			sum := 0.0
			for i := 0; i < 10; i++ {
				sum += sensorValues[i]
			}

			outputChan <- MeanResult{
				SensorValues: sensorValues,
				Result:       sum / 10.0,
			}
			sensorValues = nil
		}
	}
}
