package canvas

// Line draws a line between two points
func (ctx *Canvas) Line(x1, y1, x2, y2 int) {
	if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; y1 <= y2; y1++ {
			ctx.Pix(x1, y1)
		}
	} else if y1 == y2 {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for ; x1 <= x2; x1++ {
			ctx.Pix(x1, y1)
		}
	} else { // Bresenham
		dx := x2 - x1
		if dx < 0 {
			dx = -dx
		}
		dy := y2 - y1
		if dy < 0 {
			dy = -dy
		}
		steep := dy > dx
		if steep {
			x1, x2, y1, y2 = y1, y2, x1, x2
		}
		if x1 > x2 {
			x1, x2, y1, y2 = x2, x1, y2, y1
		}
		dx = x2 - x1
		dy = y2 - y1
		ystep := 1
		if dy < 0 {
			dy = -dy
			ystep = -1
		}
		err := dx / 2
		for ; x1 <= x2; x1++ {
			if steep {
				ctx.Pix(y1, x1)
			} else {
				ctx.Pix(x1, y1)
			}
			err -= dy
			if err < 0 {
				y1 += ystep
				err += dx
			}
		}
	}
}

/*
void Adafruit_GFX::writeFastVLine(int16_t x, int16_t y, int16_t h,
                                  uint16_t color) {
  // Overwrite in subclasses if startWrite is defined!
  // Can be just writeLine(x, y, x, y+h-1, color);
  // or writeFillRect(x, y, 1, h, color);
  drawFastVLine(x, y, h, color);
}
*/
