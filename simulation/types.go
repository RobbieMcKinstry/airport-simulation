package simulation

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

	// TODO need to track if the counter is for first class passengers or not
	Counter struct {
		State        int       // The state that the passenger is in at the current time
		IsFirstClass bool      // Represents whether or not this counter is for first class passengers or not
		current      Passenger // The person at the desk
		Time         uint64    // The time that the event is over
		A            *Airport  // A link to the parent airport
	}
)

// This block declares all of the event types.
type (
	CommuterArrival struct {
		a    *Airport
		Time uint64
	}

	InternationalArrival struct {
		a              *Airport
		ExpectedFlight *Flight
		Time           uint64
	}

	InternationalFlightTakeOff struct {
		A    *Airport
		Time uint64
	}

	CommuterFlightTakeOff struct {
		A    *Airport
		Time uint64
	}

	CheckInEmptyFirstClass struct {
		A    *Airport
		Time uint64
	}

	CheckInEmptyCoach struct {
		A    *Airport
		Time uint64
	}
)

func (c *CommuterArrival) GetTime() uint64              { return c.Time }
func (c *CommuterArrival) SetTime(t uint64)             { c.Time = t }
func (i *InternationalArrival) GetTime() uint64         { return i.Time }
func (i *InternationalArrival) SetTime(t uint64)        { i.Time = t }
func (fa *InternationalFlightTakeOff) GetTime() uint64  { return fa.Time }
func (fa *InternationalFlightTakeOff) SetTime(t uint64) { fa.Time = t }
func (cf *CommuterFlightTakeOff) GetTime() uint64       { return cf.Time }
func (cf *CommuterFlightTakeOff) SetTime(t uint64)      { cf.Time = t }
func (c *CheckInEmptyFirstClass) GetTime() uint64       { return c.Time }
func (c *CheckInEmptyFirstClass) SetTime(t uint64)      { c.Time = t }
func (c *CheckInEmptyCoach) GetTime() uint64            { return c.Time }
func (c *CheckInEmptyCoach) SetTime(t uint64)           { c.Time = t }

func (h EventHeap) Len() int           { return len(h) }
func (h EventHeap) Less(i, j int) bool { return h[i].GetTime() < h[j].GetTime() }
func (h EventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Event))
	heap.Fix(h, len(*h)-1)
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	heap.Init(h)
	return x
}
