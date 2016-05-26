package simulation

import (
	"container/heap"
	"fmt"
	. "github.com/oleiade/lane"
	"os"
)

func NewAirport() *Airport {

	eh := &EventHeap{}

	heap.Init(eh)

	a := &Airport{
		EventHeap:          eh,
		SecurityFirstClass: NewQueue(),
		SecurityCoach:      NewQueue(),
		SecurityCoach2:     NewQueue(),
		CheckInFirstClass:  []*Queue{NewQueue()},
		CheckInCoach:       []*Queue{NewQueue(), NewQueue(), NewQueue()},
		Account:            0.0,
		CommuterGate:       NewQueue(),
	}

	heap.Push(eh, &Exit{a, 1000})
	heap.Push(eh, &CommuterFlightTakeOff{a, 0})
	heap.Push(eh, &InternationalFlightTakeOff{a, 0})
	return a
}

func (a *Airport) NextEvent() Event {
	e := heap.Pop(a.EventHeap).(Event)
	return e
}

func GetShortest(qs []*Queue) *Queue {
	min := qs[0]

	for _, q := range qs {
		if q.Size() < min.Size() {
			min = q
		}
	}

	return min
}

func GetLongest(qs []*Queue) *Queue {
	max := qs[0]

	for _, q := range qs {
		if q.Size() > max.Size() {
			max = q
		}
	}

	return max
}

func (e *Exit) Visit() {
	fmt.Printf("Time: 1000 minutes\n")
	fmt.Printf("Account Balance: %v \n", e.A.Account)
	fmt.Printf("Time idle for Coach checkin: %v \n", e.A.IdleTimeCoach)
	fmt.Printf("Time idle for First Class checkin: %v \n", e.A.IdleTimeFirstClass)
	fmt.Printf("Total time idle: %v \n", e.A.IdleTime)
	os.Exit(0)
}
