package simulation

import (
	"github.com/RobbieMcKinstry/StandardNormal/stdnormal"
	rng "github.com/leesper/go_rng"

	"math"
	"math/rand"
	"time"
)

func NewExpGenerator(lambda float64) func() uint64 {
	gen := rng.NewExpGenerator(time.Now().UnixNano())

	expBuffer := make(chan float64)

	go func() {
		for {
			expBuffer <- gen.Exp(lambda)
		}
	}()

	return func() uint64 {
		return round(<-expBuffer)
	}
}

func NewGeoGenerator(p float64) func() uint64 {
	gen := rng.NewGeometricGenerator(time.Now().UnixNano())

	geoBuffer := make(chan int64)

	go func() {
		for {
			geoBuffer <- gen.Geometric(p)
		}
	}()

	return func() uint64 {
		return uint64(<-geoBuffer)
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

func NewNormalGenerator(mean, variance float64) func() uint64 {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	normBuffer := make(chan float64)

	go func() {
		for {
			normBuffer <- stdnormal.Polar(random)
		}
	}()

	return func() uint64 {
		return round(<-normBuffer)
	}

}

func NewUniformRandomGenerator() func(int) int {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	return func(n int) int {
		return random.Intn(n)
	}

}

func round(f float64) uint64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return uint64(f + math.Copysign(0.5, f))
}
