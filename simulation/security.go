package simulation

import (
	_ "container/heap"
)

func (sec *LeaveSecurityFirstClass) Visit() {

	passenger := sec.A.SecurityFirstClass.Pop().(*International)
	passenger.TakeOff.Passengers = append(passenger.TakeOff.Passengers, passenger)
}

func (sec *LeaveSecurityCoach) Visit() {
	passenger := sec.A.SecurityCoach.Pop().(*Commuter)
	sec.A.CommuterGate.Append(passenger)
}

func (sec *EmptySecurityFirstClass) Visit() {

}

func (sec *EmptySecurityCoach) Visit() {

}
