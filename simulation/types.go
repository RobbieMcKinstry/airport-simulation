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
		FirstClass  bool
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

// This block declares all of the event types.
type (
	CommuterArrival struct {
		A    *Airport
		Time uint64
	}

	InternationalArrival struct {
		A              *Airport
		ExpectedFlight *Flight
		IsFirstClass   bool
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

	TerminateProgram struct {
		A    *Airport
		Time uint64
	}

	BoardingPassPrinted struct {
		A    *Airport
		Time uint64
		Curr *Passenger
	}

	BagsChecked struct {
		A    *Airport
		Time uint64
		Curr *Passenger
	}

	MiscDelaysFinishedFC struct {
		A    *Airport
		Time uint64
		Curr *Passenger
	}

	MiscDelaysFinishedCoach struct {
		A    *Airport
		Time uint64
		Curr *Passenger
	}

	LeaveSecurityFirstClass struct {
		A    *Airport
		Time uint64
	}

	LeaveSecurityCoach struct {
		A            *Airport
		Time         uint64
		HasPassenger bool
	}

	EmptySecurityFirstClass struct {
		A    *Airport
		Time uint64
	}

	EmptySecurityCoach struct {
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
func (t *TerminateProgram) GetTime() uint64             { return t.Time }
func (t *TerminateProgram) SetTime(ti uint64)           { t.Time = ti }
func (t *BoardingPassPrinted) GetTime() uint64          { return t.Time }
func (t *BoardingPassPrinted) SetTime(ti uint64)        { t.Time = ti }
func (t *BagsChecked) GetTime() uint64                  { return t.Time }
func (t *BagsChecked) SetTime(ti uint64)                { t.Time = ti }
func (t *MiscDelaysFinishedFC) GetTime() uint64         { return t.Time }
func (t *MiscDelaysFinishedFC) SetTime(ti uint64)       { t.Time = ti }
func (t *MiscDelaysFinishedCoach) GetTime() uint64      { return t.Time }
func (t *MiscDelaysFinishedCoach) SetTime(ti uint64)    { t.Time = ti }
func (t *LeaveSecurityFirstClass) GetTime() uint64      { return t.Time }
func (t *LeaveSecurityFirstClass) SetTime(ti uint64)    { t.Time = ti }
func (t *LeaveSecurityCoach) GetTime() uint64           { return t.Time }
func (t *LeaveSecurityCoach) SetTime(ti uint64)         { t.Time = ti }
func (t *EmptySecurityCoach) GetTime() uint64           { return t.Time }
func (t *EmptySecurityCoach) SetTime(ti uint64)         { t.Time = ti }
func (t *EmptySecurityFirstClass) GetTime() uint64      { return t.Time }
func (t *EmptySecurityFirstClass) SetTime(ti uint64)    { t.Time = ti }

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
