package math_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	doodlekitMath "github.com/lucasepe/doodlekit/internal/math"
)

func TestDeg(t *testing.T) {
	table := []struct {
		in   float64
		want float64
	}{
		{1, 57.2958},
		{2, 114.5920},
		{3, 171.8873},
		{4, 229.1831},
		{5, 286.4788},
	}

	opts := cmp.Options{
		cmp.Comparer(func(x, y float64) bool {
			delta := math.Abs(x - y)
			mean := math.Abs(x+y) / 2.0
			return delta/mean < 0.0001
		}),
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("tc[%d]", i), func(t *testing.T) {
			got := doodlekitMath.Deg(tc.in)
			if !cmp.Equal(got, tc.want, opts) {
				t.Fatalf("Deg(%f) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestRad(t *testing.T) {
	table := []struct {
		in   float64
		want float64
	}{
		{57.2958, 1},
		{114.5920, 2},
		{171.8873, 3},
		{229.1831, 4},
		{286.4788, 5},
	}

	opts := cmp.Options{
		cmp.Comparer(func(x, y float64) bool {
			delta := math.Abs(x - y)
			mean := math.Abs(x+y) / 2.0
			return delta/mean < 0.0001
		}),
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("tc[%d]", i), func(t *testing.T) {
			got := doodlekitMath.Rad(tc.in)
			if !cmp.Equal(got, tc.want, opts) {
				t.Fatalf("Rad(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestTurn(t *testing.T) {
	table := []struct {
		in   float64
		want float64
	}{
		{1, 0.1591},
		{2, 0.3183},
		{3, 0.4774},
		{4, 0.6366},
		{5, 0.7957},
		{6.2831, 1},
	}

	opts := cmp.Options{
		cmp.Comparer(func(x, y float64) bool {
			delta := math.Abs(x - y)
			mean := math.Abs(x+y) / 2.0
			return delta/mean < 0.001
		}),
	}

	for i, tc := range table {
		t.Run(fmt.Sprintf("tc[%d]", i), func(t *testing.T) {
			got := doodlekitMath.Turn(tc.in)
			if !cmp.Equal(got, tc.want, opts) {
				t.Fatalf("Turn(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func ExampleSign() {

	fmt.Println(doodlekitMath.Sign(-2))
	fmt.Println(doodlekitMath.Sign(0))
	fmt.Println(doodlekitMath.Sign(2))

	// Output:
	// -1
	// 0
	// 1
}

func ExampleClamp() {
	fmt.Println(doodlekitMath.Clamp(-5, 10, 10))
	fmt.Println(doodlekitMath.Clamp(15, 10, 15))
	fmt.Println(doodlekitMath.Clamp(25, 10, 20))

	// Output:
	// 10
	// 15
	// 20
}

func ExampleLerp() {
	fmt.Println(doodlekitMath.Lerp(0, 2, 0.1))
	fmt.Println(doodlekitMath.Lerp(1, 10, 0.5))
	fmt.Println(doodlekitMath.Lerp(2, 4, 0.5))

	// Output:
	// 0.2
	// 5.5
	// 3
}

func ExampleTrunc() {
	fmt.Println(doodlekitMath.Trunc(1567.23456, 1))
	fmt.Println(doodlekitMath.Trunc(1567.23456, 0.1))
	fmt.Println(doodlekitMath.Trunc(1567.23456, 0.01))
	fmt.Println(doodlekitMath.Trunc(1567.23456, 3))

	// Output:
	// 1567
	// 1567.2
	// 1567.23
	// 1566
}
