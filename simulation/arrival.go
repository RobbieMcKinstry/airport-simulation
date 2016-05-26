package simulation

import (
	"container/heap"
	. "github.com/oleiade/lane"
)

// Make a new person and him to the queues.
// Then, add this to the heap again.
func (arr *CommuterArrival) Visit() {
	passenger := &Commuter{
		ArrivalTime: arr.Time,
		State:       QueueingForCheckIn,
		bags:        CommuterBagGen(),
	}

	shortest := GetShortest(arr.A.CheckInCoach)
	shortest.Append(passenger)

	// Add a new Arrival to the queue
	arr.Time += CommuterArrivalGen()
	arr.A.EventHeap.Push(arr)
}

// Make a new person and him to the queues.
// Then, add this to the heap again.
func (arr *InternationalArrival) Visit() {
	passenger := &International{
		ArrivalTime: arr.Time,
		State:       QueueingForCheckIn,
		bags:        InternationalBagGen(),
		FirstClass:  arr.IsFirstClass,
		TakeOff:     arr.ExpectedFlight,
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
	arr.Time += CommuterArrivalGen()
	arr.A.EventHeap.Push(arr)
}

func (fa *InternationalFlightTakeOff) Visit() {
	flight := &Flight{
		Time:                fa.Time + 6*60,
		Passengers:          make([]Passenger, 200),
		CoachSeatsFull:      0,
		FirstClassSeatsFull: 0,
		IsInternational:     true,
	}

	for i := 0; i < 50; i++ {
		if BernoulliFirstGen() {
			passenger := &InternationalArrival{
				A:              fa.A,
				ExpectedFlight: flight,
				IsFirstClass:   true,
				Time:           fa.GetTime() + InternationalArrivalGen(),
			}
			fa.A.EventHeap.Push(passenger)
		}
	}

	for i := 0; i < 150; i++ {
		if BernoulliCoachGen() {
			passenger := &InternationalArrival{
				A:              fa.A,
				ExpectedFlight: flight,
				IsFirstClass:   false,
				Time:           fa.GetTime() + InternationalArrivalGen(),
			}
			fa.A.EventHeap.Push(passenger)
		}
	}
	heap.Push(fa.A.EventHeap, &InternationalFlightTakeOff{fa.A, flight.Time})
	fa.A.Account += flight.FirstClassSeatsFull * 1000
	fa.A.Account += flight.CoachSeatsFull * 500
}

func (fa *CommuterFlightTakeOff) Visit() {
	flight := &Flight{
		Time:                fa.Time + 30,
		Passengers:          make([]Passenger, 50),
		CoachSeatsFull:      0,
		FirstClassSeatsFull: 0,
		IsInternational:     false,
	}

	for i := 0; i < 50; i++ {
		if fa.A.CommuterGate.Empty() {
			break
		}
		flight.Passengers = append(flight.Passengers, fa.A.CommuterGate.Pop().(*Commuter))
	}
	heap.Push(fa.A.EventHeap, &CommuterFlightTakeOff{fa.A, flight.Time})
	fa.A.Account += len(flight.Passengers) * 200
}
