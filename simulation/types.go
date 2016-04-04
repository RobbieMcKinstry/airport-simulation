package events

import (
	"container/heap"
	. "github.com/oleiade/lane"
)

// State machine:
// ------------------------------------------------------------------------------------
// Commuter:
// When a commuter arrives, calculate the number of bags he has.
// Printing boarding pass
// Checking bags
// Misc delays
// At securty
// At Gate
//
// International:
// When an international arrives, calculate the number of bags he has,
// Then note what time his flight is.
// Note if he's first class or coach, then he has the same state machine as a Commuter:
// Printing boarding pass
// Checking bags
// Misc delays
// At securty
// At Gate

type (
	Airport struct {
		EventHeap heap.Interface

		SecurityFirstClass *Queue
		SecurityCoach      *Queue
		SecurityCoach2     *Queue
		CheckInFirstClass  []*Queue
		CheckInCoach       []*Queue

		Account            int64
		IdleTime           uint64
		IdleTimeCoach      uint64
		IdleTimeFirstClass uint64
	}

	EventHeap []Event

	Event interface {
		Visit()
		GetTime() uint64
		SetTime(uint64)
	}

	Passenger interface {
		IsCommuter() bool
		IsFirstClass() bool
		State() int
		Bags() int64
	}

	Flight struct {
		Time                uint64
		Passengers          []*Passenger
		CoachSeatsFull      int
		FirstClassSeatsFull int
		IsInternational     bool
	}

	International struct {
		TakeOff     *Flight
		ArrivalTime uint64
		State       int
		Bags        int64
	}

	Commuter struct {
		TakeOff     *Flight
		ArrivalTime uint64
		State       int
		Bags        int64
	}
)

const (
	CommuterArrivalRate      = (40.0 / 60.0)
	InternationalArrivalMean = 75.0
	InternationalArrivalVar  = 50.0

	QueueingForCheckIn = iota
	PrintingBoardingPass
	CheckingBags
	MiscDelays
	QueueingForSecurity
	AtSecurity
	AtGate
	EmptyQueue
)

var (
	CommuterArrivalGen      = NewExpGenerator(CommuterArrivalRate)
	InternationalArrivalGen = NewNormalGenerator(InternationalArrivalMean, InternationalArrivalVar)
	BagCheckGen             = NewExpGenerator(1.0)
	BoardingPassGen         = NewExpGenerator(1.0 / 2.0)
	MiscGen                 = NewExpGenerator(1.0 / 3.0)
	CommuterBagGen          = NewGeoGenerator(0.80)
	InternationalBagGen     = NewGeoGenerator(0.60)
	BernoulliFirstGen       = NewBernGenerator(0.80)
	BernoulliCoachGen       = NewBernGenerator(0.85)
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

func (h EventHeap) Len() int           { return len(h) }
func (h EventHeap) Less(i, j int) bool { return h[i].GetTime() < h[j].GetTime() }
func (h EventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Event))
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
