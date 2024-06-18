package canvas

import (
	"image"
	"math"

	"golang.org/x/image/math/fixed"
)

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func unfixi(x fixed.Int26_6) int {
	return int(math.Round(unfix(x)))
}

func unfix(x fixed.Int26_6) float64 {
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
