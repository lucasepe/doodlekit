package canvas_test

import (
	"fmt"
	"image"
	"image/color"

	"github.com/lucasepe/doodlekit/internal/canvas"
)

func Example_canvas_Circ() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 9, 9), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Translate(4, 4)

	gc.Color(1)
	gc.Circ(0, 0, 2)

	gc.Identity()

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
	// ░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓░░░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░░░▓▓▓▓▓▓░░░░░░
	// ░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░
}

func Example_canvas_CircFill() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 9, 9), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Translate(4, 4)

	gc.Color(1)
	gc.CircFill(0, 0, 2)

	gc.Identity()

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
	// ░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓░░░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░░░▓▓▓▓▓▓░░░░░░
	// ░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░
}
