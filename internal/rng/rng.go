package rng

import (
	"math/rand"
	"time"
)

func New() RNG {
	seed := time.Now().UnixNano()
	return &rng{
		src: rand.New(rand.NewSource(seed)),
	}
}

type RNG interface {
	Seed(seed int)
	Rnd(min, max float64) float64
	RndI(min, max int) int
}

var _ RNG = (*rng)(nil)

type rng struct {
	src *rand.Rand
}

// Create and seed the generator.
// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
// Using a fixed seed will produce the same output on every run.
func (r *rng) Seed(seed int) {
	r.src = rand.New(rand.NewSource(int64(seed)))
}

func (r *rng) Rnd(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}

	return min + r.src.Float64()*(max-min)
}

func (r *rng) RndI(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return r.src.Intn(max+1-min) + min
}
