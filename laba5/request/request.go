package request

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Option 2:
//
//	F(t) = t/10;
//	G(t) =1-e^(-mt)
const (
	lambda = 0.1
	miu    = 2
)

type request struct {
	Probability  float32
	TimeToArrive float64
	TimeToServe  float64
}

type Requests []request

func (r *Requests) GenerateArrivalTimes() {
	for i := 0; i < len(*r); i++ {
		(*r)[i].TimeToArrive = float64((*r)[i].Probability) * 10
	}
}

func (r *Requests) GenerateServingTimes() {
	for i := 0; i < len(*r); i++ {
		(*r)[i].TimeToServe = float64((math.Log(float64(float32(1) - (*r)[i].Probability))) / (-miu))
	}
}

func (r Requests) GetTotalArriveTime() (ret float64) {
	for _, item := range r {
		ret += item.TimeToArrive
	}
	return
}

func (r Requests) GetTotalServingTime() (ret float64) {
	for _, item := range r {
		ret += item.TimeToServe
	}
	return
}
func (r Requests) GetTotalTime() (ret float64) {
	ret = r.GetTotalArriveTime() + r.GetTotalServingTime()
	return
}

func (r Requests) GenerateRandomSequence(seqLen int) {
	random := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < seqLen; i++ {
		r[i].Probability = float32(random.Uint32()%100) / 100
	}
}

func (r Requests) String() (ret string) {
	ret += "Requests{[\n"
	for i, item := range r {
		ret += fmt.Sprintf("\tindex - %v : { probability: %v, timeToArrive: %.3v sec, timeToServe: %.3v sec}\n",
			i+1, item.Probability, item.TimeToArrive, item.TimeToServe)
	}
	ret += "]\n"
	ret += fmt.Sprintf("totalArrivalTime : %v sec\n", r.GetTotalArriveTime())
	ret += fmt.Sprintf("totalServingTime : %v sec\n", r.GetTotalServingTime())
	ret += fmt.Sprintf("totalTime : %v sec\n", r.GetTotalTime())
	ret += fmt.Sprintf("AverageArrivalTime : %v sec\n", r.GetTotalArriveTime()/float64(len(r)))
	ret += fmt.Sprintf("AverageServiceTime : %v sec\n", r.GetTotalServingTime()/float64(len(r)))
	ret += "}\n"
	return
}
