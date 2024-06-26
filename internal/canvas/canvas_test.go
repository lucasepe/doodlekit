package canvas_test

import (
	"fmt"
	"image"
	"image/color"

	"github.com/lucasepe/doodlekit/internal/canvas"
)

func Example_canvas_Print() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0x00},
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	out := image.NewPaletted(image.Rect(0, 0, 21, 5), pal)

	gc := canvas.New(out)
	gc.Cls(1)

	gc.Color(2)
	gc.Print("Hello!", 1, 0)

	// //gc.Color(7)
	// //gc.Rect(0, 0, 3, 4)
	// gc.Save("before.gif")
	// letH := gc.Smake(0, 0, 3, 5, 12, 15)
	// gc.Reset()

	// gc.Cls(1)
	// gc.Translate(float64(gc.Width())*0.5, float64(gc.Height())*0.5)
	// gc.Rotate(0.785398)
	// gc.Sput(letH, 0, 0, 0.5, 0.5)
	// //gc.Record()
	// gc.Save("after.gif")

	for y := 0; y < out.Bounds().Dy(); y++ {
		for x := 0; x < out.Bounds().Dx(); x++ {
			if out.ColorIndexAt(x, y) == 2 {
				fmt.Print("▓▓")
			} else {
				fmt.Print("░░")
			}
		}
		fmt.Println()
	}

	// Output:
	// ░░▓▓░░▓▓░░░░░░░░░░▓▓▓▓░░░░▓▓▓▓░░░░░░░░░░░░
	// ░░▓▓░░▓▓░░▓▓▓▓▓▓░░░░▓▓░░░░░░▓▓░░░░▓▓▓▓▓▓░░
	// ░░▓▓▓▓▓▓░░▓▓░░▓▓░░░░▓▓░░░░░░▓▓░░░░▓▓░░▓▓░░
	// ░░▓▓░░▓▓░░▓▓▓▓░░░░░░▓▓░░░░░░▓▓░░░░▓▓░░▓▓░░
	// ░░▓▓░░▓▓░░▓▓▓▓▓▓░░░░▓▓░░░░░░▓▓░░░░▓▓▓▓▓▓░░
}

func Example_canvas_Smake() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0x00},
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	out := image.NewPaletted(image.Rect(0, 0, 16, 10), pal)

	gc := canvas.New(out)
	gc.Cls(1)

	gc.Color(2)
	gc.Print("Hi!", 0, 0)

	out = gc.Smake(0, 0, 3, 5, 9, 10)

	for y := 0; y < out.Bounds().Dy(); y++ {
		for x := 0; x < out.Bounds().Dx(); x++ {
			if out.ColorIndexAt(x, y) == 2 {
				fmt.Print("▓▓")
			} else {
				fmt.Print("░░")
			}
		}
		fmt.Println()
	}

	// Output:
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
	// ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
	// ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
	// ▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓
}

func Example_canvas_Sput() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0x00},
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	out := image.NewPaletted(image.Rect(0, 0, 22, 12), pal)

	gc := canvas.New(out)
	gc.Cls(1)

	gc.Color(2)
	gc.Print("Hi!", 0, 0)

	letH := gc.Smake(0, 0, 3, 5, 15, 10)
	gc.Cls(1)

	gc.Translate(10, 6)
	gc.Sput(letH, 0, 0, 0.5, 0.5)

	for y := 0; y < out.Bounds().Dy(); y++ {
		for x := 0; x < out.Bounds().Dx(); x++ {
			if out.ColorIndexAt(x, y) == 2 {
				fmt.Print("▓▓")
			} else {
				fmt.Print("░░")
			}
		}
		fmt.Println()
	}

	// Output:
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
}
