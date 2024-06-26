package canvas

import (
	"image"

	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
)

// Smake creates a sprite from the specified rect on canvas.
// The sprits can be stretched if sw and/or sh are > 0.
//
// cx: x coordinate of the upper left corner of the rectangle in the canvas.
// cy: y coordinate of the upper left corner of the rectangle in the canvas.
// cw: width of the rectangle in the canvas, as a number of pixels.
// ch: height of the rectangle in the canvas, as a number of pixels.
//
// sw: width of the sprite.
// sh: height of sprite.
func (ctx *Canvas) Smake(cx, cy, cw, ch, sw, sh int) *image.Paletted {
	if sw <= 0 {
		sw = cw
	}

	if sh <= 0 {
		sh = ch
	}

	dst := image.NewPaletted(image.Rect(0, 0, sw, sh), ctx.img.Palette)
	draw.NearestNeighbor.Scale(dst, dst.Rect,
		ctx.img, image.Rect(cx, cy, cx+cw, cy+ch),
		draw.Over, nil)

	for x := dst.Rect.Min.X; x < dst.Rect.Max.X; x++ {
		for y := dst.Rect.Min.Y; y < dst.Rect.Max.Y; y++ {
			if dst.ColorIndexAt(x, y) == ctx.bgColorIndex {
				dst.SetColorIndex(x, y, 0)
			}
		}
	}

	// f, err := os.Create("smake.png")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	// }
	// defer f.Close()
	// if err := png.Encode(f, dst); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	// }

	return dst
}

// Sput draws the specified image at the specified anchor point.
// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
// image. Use ax=0.5, ay=0.5 to center the image at the specified point.
func (ctx *Canvas) Sput(im *image.Paletted, x, y int, ax, ay float64) {
	s := im.Bounds().Size()

	x -= int(ax * float64(s.X))
	y -= int(ay * float64(s.Y))

	fx, fy := float64(x), float64(y)
	m := ctx.matrix.Translate(fx, fy)
	s2d := f64.Aff3{m.XX, m.XY, m.X0, m.YX, m.YY, m.Y0}

	transformer := draw.BiLinear
	transformer.Transform(ctx.img, s2d, im, im.Bounds(), draw.Over, nil)
}
