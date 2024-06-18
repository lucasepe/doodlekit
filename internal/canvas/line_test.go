package canvas_test

import (
	"fmt"
	"image"
	"image/color"

	"github.com/lucasepe/doodlekit/internal/canvas"
)

func Example_canvas_Line() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 8, 8), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Color(1)

	gc.Line(1, 1, 6, 6)

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			if dst.ColorIndexAt(x, y) == 1 {
				fmt.Print("▓▓")
			} else {
				fmt.Print("░░")
			}
		}
		fmt.Println()
	}

	// Output:
	// ░░░░░░░░░░░░░░░░
	// ░░▓▓░░░░░░░░░░░░
	// ░░░░▓▓░░░░░░░░░░
	// ░░░░░░▓▓░░░░░░░░
	// ░░░░░░░░▓▓░░░░░░
	// ░░░░░░░░░░▓▓░░░░
	// ░░░░░░░░░░░░▓▓░░
	// ░░░░░░░░░░░░░░░░
}

func Example_canvas_Line_H() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 8, 3), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Color(1)
	gc.Line(1, 1, 6, 1)

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			if dst.ColorIndexAt(x, y) == 1 {
				fmt.Print("▓▓")
			} else {
				fmt.Print("░░")
			}
		}
		fmt.Println()
	}

	// Output:
	// ░░░░░░░░░░░░░░░░
	// ░░▓▓▓▓▓▓▓▓▓▓▓▓░░
	// ░░░░░░░░░░░░░░░░
}

func Example_canvas_Line_V() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 3, 8), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Color(1)
	gc.Line(1, 1, 1, 6)

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			if dst.ColorIndexAt(x, y) == 1 {
				fmt.Print("▓▓")
			} else {
				fmt.Print("░░")
			}
		}
		fmt.Println()
	}

	// Output:
	// ░░░░░░
	// ░░▓▓░░
	// ░░▓▓░░
	// ░░▓▓░░
	// ░░▓▓░░
	// ░░▓▓░░
	// ░░▓▓░░
	// ░░░░░░
}
