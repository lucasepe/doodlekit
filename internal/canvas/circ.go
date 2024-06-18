package canvas

// Circ draws a circle given a point and radius
func (ctx *Canvas) Circ(x0, y0, r int) {
	f := 1 - r
	ddfx := 1
	ddfy := -2 * r
	x := 0
	y := r

	ctx.Pix(x0, y0+r)
	ctx.Pix(x0, y0-r)
	ctx.Pix(x0+r, y0)
	ctx.Pix(x0-r, y0)

	for x < y {
		if f >= 0 {
			y--
			ddfy += 2
			f += ddfy
		}
		x++
		ddfx += 2
		f += ddfx

		ctx.Pix(x0+x, y0+y)
		ctx.Pix(x0-x, y0+y)
		ctx.Pix(x0+x, y0-y)
		ctx.Pix(x0-x, y0-y)
		ctx.Pix(x0+y, y0+x)
		ctx.Pix(x0-y, y0+x)
		ctx.Pix(x0+y, y0-x)
		ctx.Pix(x0-y, y0-x)
	}
}

func (ctx *Canvas) CircFill(x, y, r int) {
	ctx.Line(x, y-r, x, y+r)
	ctx.fillCircleHelper(x, y, r, 3, 0)
}

func (ctx *Canvas) drawCircleHelper(x0, y0, r int, cornername uint8) {
	f := (1 - r)
	ddfx := 1
	ddfy := -2 * r
	x := 0
	y := r

	for x < y {
		if f >= 0 {
			y--
			ddfy += 2
			f += ddfy
		}
		x++
		ddfx += 2
		f += ddfx

		if cornername&0x4 != 0 {
			ctx.Pix(x0+x, y0+y)
			ctx.Pix(x0+y, y0+x)
		}
		if cornername&0x2 != 0 {
			ctx.Pix(x0+x, y0-y)
			ctx.Pix(x0+y, y0-x)
		}
		if cornername&0x8 != 0 {
			ctx.Pix(x0-y, y0+x)
			ctx.Pix(x0-x, y0+y)
		}
		if cornername&0x1 != 0 {
			ctx.Pix(x0-y, y0-x)
			ctx.Pix(x0-x, y0-y)
		}
	}
}

// Quarter-circle drawer with fill, used for circles and roundrects
// x0       Center-point x coordinate
// y0       Center-point y coordinate
// r        Radius of circle
// corners  Mask bits indicating which quarters we're doing
// delta    Offset from center-point, used for round-rects
func (ctx *Canvas) fillCircleHelper(x0, y0, r int, corners uint8, delta int) {
	f := 1 - r
	ddF_x := 1
	ddF_y := -2 * r
	x := 0
	y := r
	px := x
	py := y

	delta++ // Avoid some +1's in the loop

	for x < y {
		if f >= 0 {
			y--
			ddF_y += 2
			f += ddF_y
		}
		x++
		ddF_x += 2
		f += ddF_x

		if x < (y + 1) {
			if corners&1 != 0 {
				// writeFastVLine(x0 + x, y0 - y, 2 * y + delta, color);
				// writeLine(x, y, x, y+h-1, color);

				ctx.Line(x0+x, y0-y, x0+x, y0+y+delta-1)
			}
			if corners&2 != 0 {
				// writeFastVLine(x0 - x, y0 - y, 2 * y + delta, color);
				ctx.Line(x0-x, y0-y, x0-x, y0+y+delta-1)
			}
		}
		if y != py {
			if corners&1 != 0 {
				// writeFastVLine(x0 + py, y0 - px, 2 * px + delta, color);
				ctx.Line(x0+py, y0-px, x0+py, y0+px+delta-1)
			}
			if corners&2 != 0 {
				// writeFastVLine(x0 - py, y0 - px, 2 * px + delta, color);
				ctx.Line(x0-py, y0-px, x0-py, y0+px+delta-1)
			}
			py = y
		}
		px = x
	}
}
