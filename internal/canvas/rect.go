package canvas

// Rect draws a rectangle.

/*
(x1,y1)          (x2,y1)

	+----------------+
	|                |
	|                |
	+----------------+

(x1,y2)          (x2,y2)
*/
func (ctx *Canvas) Rect(x1, y1, x2, y2 int) {
	ctx.Line(x1, y1, x2, y1)
	ctx.Line(x2, y1, x2, y2)
	ctx.Line(x2, y2, x1, y2)
	ctx.Line(x1, y2, x1, y1)
}

// RectFill draws a filled rectangle given a point, width and height
func (ctx *Canvas) RectFill(x1, y1, x2, y2 int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	for i := x1; i <= x2; i++ {
		ctx.Line(i, y1, i, y2)
	}
}

func (ctx *Canvas) RoundRect(x1, y1, x2, y2, r int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	w, h := (x2 - x1), (y2 - y1)
	maxRadius := w / 2
	if w < h {
		maxRadius = h / 2
	}
	if r > maxRadius {
		r = maxRadius
	}

	ctx.Line(x1+r, y1, x2-r, y1) // Top
	ctx.Line(x1+r, y2, x2-r, y2) // Bottom
	ctx.Line(x1, y1+r, x1, y2-r) // Left
	ctx.Line(x2, y1+r, x2, y2-r) // Right

	// Draw four corners
	ctx.drawCircleHelper(x1+r+1, y1+r, r, uint8(1))
	ctx.drawCircleHelper(x2-r-1, y1+r, r, uint8(2))
	ctx.drawCircleHelper(x2-r-1, y2-r, r, uint8(4))
	ctx.drawCircleHelper(x1+r, y2-r-1, r, uint8(8))
}

func (ctx *Canvas) RoundRectFill(x1, y1, x2, y2, r int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	w, h := (x2 - x1), (y2 - y1)
	maxRadius := w / 2
	if w < h {
		maxRadius = h / 2
	}
	if r > maxRadius {
		r = maxRadius
	}

	for i := x1 + r; i <= x2-r; i++ {
		ctx.Line(i, y1, i, y2)
	}

	// Draw the four rounded corners
	// fillCircleHelper(x + w - r - 1, y + r, r, 1, h - 2 * r - 1, color);
	// fillCircleHelper(x + r, y + r, r, 2, h - 2 * r - 1, color);
	ctx.fillCircleHelper(x2-r, y1+r+1, r, uint8(1), h-2*r-2)
	ctx.fillCircleHelper(x1+r, y1+r+1, r, uint8(2), h-2*r-2)
}
