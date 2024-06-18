package math

import "math"

func Deg(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}

func Rad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func Turn(rad float64) float64 {
	return rad * 0.159155
}

// Sign returns -1 for values < 0, 0 for 0, and 1 for values > 0.
func Sign(x float64) float64 {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

// Clamp returns x clamped to the interval [min, max].
//
// If x is less than min, min is returned. If x is more than max, max is returned. Otherwise, x is
// returned.
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Lerp does linear interpolation between two values.
func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

// Trunc throws away the fraction digits
//
// in := 1567.23456
// out := Trunc(in, 1)
// fmt.Printf("%.6f\n", out) -> 1567.000000
//
// in = 1.234567
// out = Trunc(in, 0.01)
// fmt.Printf("%.6f\n", out) -> 1.230000
func Trunc(f float64, unit float64) float64 {
	return math.Trunc(f/unit) * unit
}
