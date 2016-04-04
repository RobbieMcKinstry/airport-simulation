package simulation

import (
	"container/heap"
	. "github.com/oleiade/lane"
)

func NewAirport() *Airport {

	eh := &EventHeap{}
	heap.Init(eh)

	return &Airport{
		EventHeap:          eh,
		SecurityFirstClass: NewQueue(),
		SecurityCoach:      NewQueue(),
		SecurityCoach2:     NewQueue(),
		CheckInFirstClass:  []*Queue{NewQueue()},
		CheckInCoach:       []*Queue{NewQueue(), NewQueue(), NewQueue()},
		Account:            0.0,
	}
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
