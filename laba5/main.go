package main

import (
	"fmt"

	"github.com/Atomicall/Mod/laba5/drawing"
	"github.com/Atomicall/Mod/laba5/request"
)

// Option 2:
//
//	F(t) = t/10;
//	G(t) =1-e^(-mt)

var (
	seqLen   int     = 10
	Lambda   float32 = 0.1
	Miu      int     = 2
	savePath         = "./output.png"
)

func main() {
	sequence := make(request.Requests, seqLen)
	sequence.GenerateRandomSequence(seqLen)
	sequence.GenerateArrivalTimes()
	sequence.GenerateServingTimes()
	drawing.DrawGraphic(sequence, savePath)
	fmt.Println(sequence)
}
