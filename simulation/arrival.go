package simulation

import (
	. "github.com/oleiade/lane"
)

// Make a new person and him to the queues.
// Then, add this to the heap again.
func (arr *CommuterArrival) Visit() {
	passenger := &Commuter{
		ArrivalTime: arr.Time,
		State:       QueueingForCheckIn,
		Bags:        CommuterBagGen(),
	}

	shortest := GetShortest(arr.a.CheckInCoach)
	shortest.Append(passenger)

	// Add a new Arrival to the queue
	arr.Time += uint64(round(CommuterArrivalGen()))
	arr.a.EventHeap.Push(arr)
}

// Make a new person and him to the queues.
// Then, add this to the heap again.
func (arr *InternationalArrival) Visit() {
	passenger := &International{
		ArrivalTime: arr.Time,
		State:       QueueingForCheckIn,
		Bags:        InternationalBagGen(),
		FirstClass:  arr.IsFirstClass,
	}

	// Add the passenger to the shortest line available for him
	var shortest *Queue
	if arr.IsFirstClass {
		shortest = GetShortest(arr.A.CheckInFirstClass)
	} else {
		shortest = GetShortest(arr.A.CheckInCoach)
	}
	shortest.Append(passenger)

	// Add the next arrival to the heap
	arr.Time += uint64(round(CommuterArrivalGen()))
	arr.A.EventHeap.Push(arr)
}

func (fa *InternationalFlightTakeOff) Visit() {
	flight := &Flight{
		Time:                fa.Time + 6*60,
		Passengers:          make([]*Passenger, 200),
		CoachSeatsFull:      0,
		FirstClassSeatsFull: 0,
		IsInternational:     true,
	}
	fa.A.EventHeap.Push(flight)

	for i := 0; i < 50; i++ {
		if BernoulliFirstGen() {
			passenger := InternationalArrival{
				A:              fa.A,
				ExpectedFlight: flight,
				IsFirstClass:   true,
				Time:           fa.GetTime() + uint64(round(InternationalArrivalGen())),
			}
			fa.A.EventHeap.Push(passenger)
		}
	}

	for i := 0; i < 150; i++ {
		if BernoulliCoachGen() {
			passenger := InternationalArrival{
				A:              fa.A,
				ExpectedFlight: flight,
				IsFirstClass:   false,
				Time:           fa.GetTime() + uint64(round(InternationalArrivalGen())),
			}
			fa.A.EventHeap.Push(passenger)
		}
	}
}

func (fa *CommuterFlightTakeOff) Visit() {
	flight := &Flight{
		Time:                fa.Time + 30,
		Passengers:          make([]*Passenger, 50),
		CoachSeatsFull:      0,
		FirstClassSeatsFull: 0,
		IsInternational:     false,
	}
	fa.A.EventHeap.Push(flight)

	// TODO need to pull the commuters out of the gate queue
}
