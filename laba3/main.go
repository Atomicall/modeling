package main

import (
	"fmt"

	"github.com/Atomicall/Mod/laba3/packages/stdoutlog"
)

const (
	Iterations int     = 200
	A          int     = 131
	C          int     = 1021
	M          int     = 100
	initProb   float64 = 0.9
)

func ComputeNextRandom(prevVal float64) (nextVal float64) {
	nextVal = float64(int(A*int(prevVal*100)+C)%int(M)) / 100.0
	return
}

func ComputeNextState(ConversionMat [3][3]float64, prevState int, currentProbability float64) (nextState int) {
	conversions := ConversionMat[prevState]
	accum := 0.0
	for index, convProb := range conversions {
		if (accum) < currentProbability && currentProbability < (accum+convProb) {
			nextState = index
			break
		}
		accum += convProb
	}
	return
}

func ImmitateConversions(ConversionMat [3][3]float64, initState int) (steadyProbs [3]float64, steps int) {
	stateCounter := make([]int, 3)
	steps = 1

	tempVal := ComputeNextRandom(initProb)
	nextState := ComputeNextState(ConversionMat, initState, tempVal)
	stateCounter[nextState]++

	for i := 0; i < Iterations-1; i++ {
		tempVal = ComputeNextRandom(tempVal)
		nextState = ComputeNextState(ConversionMat, nextState, tempVal)
		if ConversionMat[nextState][nextState] == 1 {
			fmt.Printf("\n!!!Got into Absorbing state!!!")
			return
		}
		stateCounter[nextState]++
		steps++
	}
	fmt.Printf("\n\t%v", stateCounter)
	for i, counts := range stateCounter {
		steadyProbs[i] = float64(counts) / float64(Iterations)
	}

	return
}

func main() {
	var (
		ConversionMat = [3][3]float64{
			{0.1, 0.8, 0.1},
			{0.3, 0.5, 0.2},
			{0.7, 0.2, 0.1},
		}
		ConversionAbsorbingMat = [3][3]float64{
			{0.2, 0.1, 0.7},
			{0, 1, 0},
			{0.3, 0.6, 0.1},
		}
	)
	ImmitateConversionsWithOutput := stdoutlog.ShowDecorator(ImmitateConversions)
	for i := 0; i < 3; i++ {
		fmt.Printf("Init State : %v\n", i)
		ImmitateConversionsWithOutput(ConversionMat, i)
		ImmitateConversionsWithOutput(ConversionAbsorbingMat, i)
		fmt.Println()
	}
}
