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
	Rnd(max float64) float64
	RndI(max int) int
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

func (r *rng) Rnd(max float64) float64 {
	return r.src.Float64() * max
}

func (r *rng) RndI(max int) int {
	if max <= 0 {
		return 0
	}
	return r.src.Intn(max + 1)
}
