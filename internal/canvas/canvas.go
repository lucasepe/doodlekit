package canvas

import (
	"image"
	"image/gif"
	"os"

	"github.com/lucasepe/doodlekit/internal/fonts"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/f64"
	"golang.org/x/image/math/fixed"
)

const (
	defaultWidth  = 64
	defaultHeight = 64
	defaultDelay  = 5
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

type Canvas struct {
	img          *image.Paletted
	fgColorIndex uint8
	bgColorIndex uint8
	resize       int
	fontFace     font.Face
	anim         *gif.GIF
	matrix       Matrix
}

func New(img *image.Paletted) *Canvas {
	res := &Canvas{
		img:      img,
		anim:     &gif.GIF{},
		fontFace: MustLoadFont(fonts.Micro()).NewFace(),
		resize:   1,
		matrix:   Identity(),
	}

	if res.img == nil {
		res.img = image.NewPaletted(
			image.Rect(0, 0, defaultWidth, defaultHeight),
			doodleKitPalette(),
		)
	}

	res.Cls(7)
	res.Color(0)

	return res
}

func (ctx *Canvas) Resize(sf int) {
	if sf <= 0 {
		sf = 1
	}

	if sf > 10 {
		sf = 10
	}

	ctx.resize = sf
}

func (ctx *Canvas) Width() int { return ctx.img.Rect.Dx() }

func (ctx *Canvas) Height() int { return ctx.img.Rect.Dy() }

func (ctx *Canvas) Color(idx int) {
	if idx < 0 {
		ctx.fgColorIndex = uint8(0)
		return
	}

	ctx.fgColorIndex = uint8(idx % len(ctx.img.Palette))
}

func (ctx *Canvas) Cls(idx int) {
	if idx < 0 {
		ctx.bgColorIndex = uint8(0)
	} else {
		ctx.bgColorIndex = uint8(idx % len(ctx.img.Palette))
	}

	draw.Draw(
		ctx.img,
		ctx.img.Bounds(),
		image.NewUniform(ctx.img.Palette[ctx.bgColorIndex]),
		image.Point{},
		draw.Src,
	)
}

func (ctx *Canvas) MeasureString(s string) (int, int) {
	d := &font.Drawer{
		Face: ctx.fontFace,
	}
	tw := d.MeasureString(s).Ceil()
	th := UnfixI(ctx.fontFace.Metrics().Height)
	return tw, th
}

func (ctx *Canvas) Print(s string, x, y int) {
	ax, ay := 0.0, -1.0
	w, h := ctx.MeasureString(s)
	x1 := float64(x) - ax*float64(w)
	y1 := float64(y) - ay*float64(h)
	ctx.drawString(s, x1, y1)
}

func (ctx *Canvas) drawString(s string, x, y float64) {
	d := &font.Drawer{
		Dst:  ctx.img,
		Src:  image.NewUniform(ctx.img.Palette[ctx.fgColorIndex]),
		Face: ctx.fontFace,
		Dot: fixed.Point26_6{
			X: Fix(x),
			Y: Fix(y),
		},
	}

	// based on Drawer.DrawString() in golang.org/x/image/font/font.go
	prevC := rune(-1)
	for _, c := range s {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			// TODO: is falling back on the U+FFFD glyph the responsibility of
			// the Drawer or the Face?
			// TODO: set prevC = '\ufffd'?
			continue
		}
		sr := dr.Sub(dr.Min)
		transformer := draw.BiLinear
		fx, fy := float64(dr.Min.X), float64(dr.Min.Y)
		m := ctx.matrix.Translate(fx, fy)
		s2d := f64.Aff3{m.XX, m.XY, m.X0, m.YX, m.YY, m.Y0}
		transformer.Transform(d.Dst, s2d, d.Src, sr, draw.Over, &draw.Options{
			SrcMask:  mask,
			SrcMaskP: maskp,
		})
		d.Dot.X += advance
		prevC = c
	}
}

func (ctx *Canvas) Pix(x, y int) {
	tx, ty := ctx.TransformPoint(float64(x), float64(y))
	ctx.img.SetColorIndex(int(tx), int(ty), ctx.fgColorIndex)
}

func (ctx *Canvas) At(x, y int) int {
	tx, ty := ctx.TransformPoint(float64(x), float64(y))
	c := ctx.img.ColorIndexAt(int(tx), int(ty))
	return int(c)
}

func (ctx *Canvas) Record() {
	var dst *image.Paletted
	if ctx.resize > 1 {
		dst = scaledImage(ctx.img, ctx.resize)
	} else {
		dst = image.NewPaletted(ctx.img.Bounds(), ctx.img.Palette)
		draw.Draw(dst, dst.Bounds(), ctx.img, image.Point{}, draw.Src)
	}

	ctx.anim.Image = append(ctx.anim.Image, dst)
	ctx.anim.Delay = append(ctx.anim.Delay, defaultDelay)
	ctx.anim.Disposal = append(ctx.anim.Disposal, gif.DisposalBackground)
}

func (ctx *Canvas) Reset() {
	ctx.anim.Image = nil
	ctx.anim.Delay = nil
	ctx.anim.Disposal = nil
}

func (ctx *Canvas) Save(fn string) error {
	if len(ctx.anim.Image) == 0 {
		ctx.Record()
	}

	w, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer w.Close()

	return gif.EncodeAll(w, ctx.anim)
}

// Transformation Matrix Operations

// Identity resets the current transformation matrix to the identity matrix.
// This results in no translating, scaling, rotating, or shearing.
func (dc *Canvas) Identity() {
	dc.matrix = Identity()
}

// Translate updates the current matrix with a translation.
func (dc *Canvas) Translate(x, y float64) {
	dc.matrix = dc.matrix.Translate(x, y)
}

// Scale updates the current matrix with a scaling factor.
// Scaling occurs about the origin.
func (dc *Canvas) Scale(x, y float64) {
	dc.matrix = dc.matrix.Scale(x, y)
}

// ScaleAbout updates the current matrix with a scaling factor.
// Scaling occurs about the specified point.
func (dc *Canvas) ScaleAbout(sx, sy, x, y float64) {
	dc.Translate(x, y)
	dc.Scale(sx, sy)
	dc.Translate(-x, -y)
}

// Rotate updates the current matrix with a anticlockwise rotation.
// Rotation occurs about the origin. Angle is specified in radians.
func (dc *Canvas) Rotate(angle float64) {
	dc.matrix = dc.matrix.Rotate(angle)
}

// RotateAbout updates the current matrix with a anticlockwise rotation.
// Rotation occurs about the specified point. Angle is specified in radians.
func (dc *Canvas) RotateAbout(angle, x, y float64) {
	dc.Translate(x, y)
	dc.Rotate(angle)
	dc.Translate(-x, -y)
}

// Shear updates the current matrix with a shearing angle.
// Shearing occurs about the origin.
func (dc *Canvas) Shear(x, y float64) {
	dc.matrix = dc.matrix.Shear(x, y)
}

// ShearAbout updates the current matrix with a shearing angle.
// Shearing occurs about the specified point.
func (dc *Canvas) ShearAbout(sx, sy, x, y float64) {
	dc.Translate(x, y)
	dc.Shear(sx, sy)
	dc.Translate(-x, -y)
}

// TransformPoint multiplies the specified point by the current matrix,
// returning a transformed position.
func (dc *Canvas) TransformPoint(x, y float64) (tx, ty float64) {
	return dc.matrix.TransformPoint(x, y)
}

// InvertY flips the Y axis so that Y grows from bottom to top and Y=0 is at
// the bottom of the image.
func (dc *Canvas) InvertY() {
	dc.Translate(0, float64(dc.Height()))
	dc.Scale(1, -1)
}
