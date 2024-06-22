package canvas

import (
	"image"
	"math"
	"strconv"

	"golang.org/x/image/math/fixed"
)

// EachPixel calls the provided function for each pixel in the provided rectangle.
func EachPixel(r image.Rectangle, fn func(x, y int)) {
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			fn(x, y)
		}
	}
}

// Converts the given float64 to the nearest fixed.Int26_6.
// If there's a tie, returned values will be different, and
// the first will always be smaller than the second.
//
// The function will panic if the given float64 is not closely
// representable by any fixed.Int26_6 (including Inf, -Inf and NaN).
func FromFloat64(value float64) (fixed.Int26_6, fixed.Int26_6) {
	// TODO: overflows may still be possible, and faster conversion
	//       methods must exist, but go figure
	candidateA := fixed.Int26_6(value * 64)
	diffA := abs64(float64(candidateA)/64.0 - value)
	if diffA == 0 {
		return candidateA, candidateA
	} // fast exact conversion

	// fast path didn't succeed, proceed now to the more complex cases

	// check NaN
	if math.IsNaN(value) {
		panic("can't convert NaN to fixed.Int26_6")
	}

	// check bounds
	if value > 33554431.984375 {
		if value <= 33554432 {
			result := fixed.Int26_6(0x7FFFFFFF)
			return result, result
		}
		given := strconv.FormatFloat(value, 'f', -1, 64)
		panic("can't convert " + given + " to fixed.Int26_6, the biggest representable value is 33554431.984375")
	} else if value < -33554432 {
		if value >= -33554432.015625 {
			result := -fixed.Int26_6(0x7FFFFFFF) - 1
			return result, result
		}
		given := strconv.FormatFloat(value, 'f', -1, 64)
		panic("can't convert " + given + " to fixed.Int26_6, the smallest representable value is -33554432.0")
	}

	// compare current candidate with the next and previous ones
	candidateB := candidateA + 1
	candidateC := candidateA - 1
	diffB := abs64(float64(candidateB)/64.0 - value)
	diffC := abs64(float64(candidateC)/64.0 - value)

	if diffA < diffB {
		if diffA == diffC {
			return candidateC, candidateA
		}
		if diffA < diffC {
			return candidateA, candidateA
		}
		return candidateC, candidateC
	} else if diffB < diffA {
		if diffB == diffC {
			panic(value)
		} // this shouldn't be possible, but just to be safe
		if diffB < diffC {
			return candidateB, candidateB
		}
		return candidateC, candidateC
	} else { // diffA == diffB
		return candidateA, candidateB
	}
}

func Fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func UnfixI(x fixed.Int26_6) int {
	return int(math.Round(Unfix(x)))
}

func Unfix(x fixed.Int26_6) float64 {
	const shift, mask = 6, 1<<6 - 1
	if x >= 0 {
		return float64(x>>shift) + float64(x&mask)/64
	}

	x = -x
	if x >= 0 {
		return -(float64(x>>shift) + float64(x&mask)/64)
	}

	return 0
}

// Doesn't care about NaNs and general floating point quirkiness.
func abs64(value float64) float64 {
	if value >= 0 {
		return value
	}
	return -value
}

// scaledImage resize using nearest neighbor scaling on dst from src.
func scaledImage(src *image.Paletted, s int) *image.Paletted {
	b := src.Bounds()
	if b.Empty() {
		return &image.Paletted{}
	}

	dst := image.NewPaletted(image.Rect(0, 0, b.Dx()*s, b.Dy()*s), src.Palette)

	w := dst.Bounds().Dx()
	h := dst.Bounds().Dy()

	xRatio := b.Dx()<<16/w + 1
	yRatio := b.Dy()<<16/h + 1

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			sx := ((x * xRatio) >> 16)
			sy := ((y * yRatio) >> 16)

			dst.Set(x, y, src.At(sx, sy))
		}
	}

	return dst
}
