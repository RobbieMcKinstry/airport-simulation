package simulation

import (
	"github.com/RobbieMcKinstry/StandardNormal/stdnormal"
	rng "github.com/leesper/go_rng"

	"math"
	"math/rand"
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

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	normBuffer := make(chan float64)

	go func() {
		for {
			normBuffer <- stdnormal.Polar(random)
		}
	}()

	return func() float64 {
		return <-normBuffer
	}

}

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}
