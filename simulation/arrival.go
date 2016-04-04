package events

import (
	"math"
)

type CommuterArrival struct {
	a    *Airport
	time uint64
}

func (arr *CommuterArrival) Time() uint64 {
	return arr.time
}

// Make a new person and him to the queues.
// Then, add this to the heap again.
func (arr *CommuterArrival) Visit() {
	passenger := &Commuter{
		ArrivalTime: arr.time,
		State:       QueueingForCheckIn,
		Bags:        CommuterBagGen(),
	}

	shortest := GetShortest(arr.a.CheckInCoach)
	shortest.Append(passenger)

	arr.time += uint64(round(CommuterArrivalGen()))
	arr.a.EventHeap.Push(arr)
}

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

type InternationalArrival struct {
	a              *Airport
	ExpectedFlight *Flight
	time           uint64
}

func (arr *InternationalArrival) Time() uint64 {
	return arr.time
}

// Make a new person and him to the queues.
// Then, add this to the heap again.
func (arr *InternationalArrival) Visit() {
	passenger := &International{
		ArrivalTime: arr.time,
		State:       QueueingForCheckIn,
		Bags:        InternationalBagGen(),
	}

	shortest := GetShortest(arr.a.CheckInCoach)
	shortest.Append(passenger)

	arr.time += uint64(round(CommuterArrivalGen()))
	arr.a.EventHeap.Push(arr)
}

type FlightArrival struct {
	A    *Airport
	Time uint64
}

func (fa *FlightArrival) GetTime() uint64 {
	return fa.Time
}

func (fa *FlightArrival) SetTime(u uint64) {
	fa.Time = u
}

func (fa *FlightArrival) Visit() {
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
