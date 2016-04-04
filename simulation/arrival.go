package simulation

import (
	"math"
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

	arr.Time += uint64(round(CommuterArrivalGen()))
	arr.a.EventHeap.Push(arr)
}

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

// Make a new person and him to the queues.
// Then, add this to the heap again.
func (arr *InternationalArrival) Visit() {
	passenger := &International{
		ArrivalTime: arr.Time,
		State:       QueueingForCheckIn,
		Bags:        InternationalBagGen(),
	}

	shortest := GetShortest(arr.a.CheckInCoach)
	shortest.Append(passenger)

	arr.Time += uint64(round(CommuterArrivalGen()))
	arr.a.EventHeap.Push(arr)
}

func (fa *InternationalFlightTakeOff) Visit() {
	flight := &Flight{
		Time:                fa.Time + 6*60,
		Passengers:          make([]*Passenger, 200),
		CoachSeatsFull:      0,
		FirstClassSeatsFull: 0,
		IsInternational:     true,
	}

	for i := 0; i < 50; i++ {
		if BernoulliFirstGen() {
			// TODO See if we should make a new international first class passenger and queue him
			// new passenger
		}

	}

	for i := 0; i < 150; i++ {
		if BernoulliCoachGen() {
			// TODO See if we should make a new international first class passenger and queue him
			// new passenger
		}

	}

	_ = flight
	// TODO
	// make a new flight, and add it to the heap
	// This makes a ton of new passengers
	// make a new flight arrival for the same time, and add it to the heap
}
