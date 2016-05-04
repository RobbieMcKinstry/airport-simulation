package simulation

import (
	"container/heap"
)

func (sec *LeaveSecurityFirstClass) Visit() {
	// Assert that there's someone in the queue
	if sec.SecurityFirstClass.Empty() {
		panic("Empty security line, but it's supposed to have a passenger")
	}

	passenger := sec.SecurityFirstClass.Pop().(Passenger)
}

func (sec *LeaveSecurityCoach) Visit() {

}

func (sec *EmptySecurityFirstClass) Visit() {

}

func (sec *EmptySecurityCoach) Visit() {

}
