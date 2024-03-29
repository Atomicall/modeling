package stdoutLog

import (
	"fmt"
	"time"
)

type Object func(ConversionMat [][]float64, initState int, iterations *int) ([3]float64, int)

func ShowDecorator(fn Object) Object {
	return func(ConversionMat [][]float64, initState int, iterations *int) ([3]float64, int) {
		fmt.Printf("Starting immitating with:\n\tconversion matrix : %v\n\tinit State : %v", ConversionMat, initState)
		start := time.Now()
		SteadyProbs, Steps := fn(ConversionMat, initState, iterations)
		fmt.Printf("\nSimulation complete!\n\tSteps taken : %v\n\tSteady Probabilities: %v\n\tTime spent : %v\n", Steps, SteadyProbs, time.Since(start))
		return SteadyProbs, Steps
	}
}
