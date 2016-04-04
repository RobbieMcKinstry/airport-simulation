package spec

import (
	sim "github.com/RobbieMcKinstry/airport-simulation/simulation"

	"container/heap"
	"testing"
)

func TestHeapRead(t *testing.T) {
	airport := sim.NewAirport()

	arrival1 := &sim.FlightArrival{
		A:    airport,
		Time: 7,
	}

	arrival2 := &sim.FlightArrival{
		A:    airport,
		Time: 5,
	}
	heap.Init(airport.EventHeap)
	airport.EventHeap.Push(arrival1)
	airport.EventHeap.Push(arrival2)

	if e := airport.NextEvent(); e.GetTime() != 5 {
		t.Errorf("Expected time 5, got time %v", e.GetTime())
	}

	if e := airport.NextEvent(); e.GetTime() != 7 {
		t.Errorf("Expected time 7, got time %v", e.GetTime())
	}
}
