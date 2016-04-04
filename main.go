package main

import (
	. "github.com/RobbieMcKinstry/airport-simulation"
)

func main() {
	airport := NewAirport()

	for {
		airport.NextEvent().Visit()
	}
}
