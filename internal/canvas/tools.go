package canvas

import (
	"image"
	"math"

	"golang.org/x/image/math/fixed"
)

// EachPixel calls the provided function for each pixel in the provided rectangle.
// func EachPixel(r image.Rectangle, fn func(x, y int)) {
// 	for x := r.Min.X; x < r.Max.X; x++ {
// 		for y := r.Min.Y; y < r.Max.Y; y++ {
// 			fn(x, y)
// 		}
// 	}
// }

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

// func runeToBytes(r rune) []byte {
// 	size := utf8.RuneLen(r)
// 	bs := make([]byte, size)
// 	utf8.EncodeRune(bs[0:], r)

// 	return bs
// }

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
