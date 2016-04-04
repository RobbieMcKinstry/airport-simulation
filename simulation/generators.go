package events

import (
	rng "github.com/leesper/go_rng"
	"time"
)

func NewExpGenerator(lambda float64) func() float64 {
	gen := rng.NewExpGenerator(time.Now().UnixNano())

	expBuffer := make(chan float64)

	go func() {
		for {
			expBuffer <- gen.Exp(lambda)
		}
	}()

	return func() float64 {
		return <-expBuffer
	}
}

func NewGeoGenerator(p float64) func() int64 {
	gen := rng.NewGeometricGenerator(time.Now().UnixNano())

	geoBuffer := make(chan int64)

	go func() {
		for {
			geoBuffer <- gen.Geometric(p)
		}
	}()

	return func() int64 {
		return <-geoBuffer
	}
}

func NewBernGenerator(p float64) func() bool {

	gen := rng.NewBernoulliGenerator(time.Now().UnixNano())
	bernBuffer := make(chan bool)

	go func() {
		for {
			bernBuffer <- gen.Bernoulli_P(p)
		}
	}()

	return func() bool {
		return <-bernBuffer
	}
}

func NewNormalGenerator(mean, variance float64) func() float64 {
	// TODO fix this to use my package
	return func() float64 {
		return 0.0
	}

}
