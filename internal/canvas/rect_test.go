package canvas_test

import (
	"fmt"
	"image"
	"image/color"

	"github.com/lucasepe/doodlekit/internal/canvas"
)

func Example_gfx_Rect() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 9, 11), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Translate(4, 5)

	gc.Color(1)
	gc.Rect(-2, -4, 2, 4)
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
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░░░░░░░░░░░░░░░
}

func Example_gfx_Rect_2() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 11, 7), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Translate(5, 3)

	gc.Color(1)
	gc.Rect(-4, -2, 4, 2)

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
	// ░░░░░░░░░░░░░░░░░░░░░░
	// ░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░
	// ░░▓▓░░░░░░░░░░░░░░▓▓░░
	// ░░▓▓░░░░░░░░░░░░░░▓▓░░
	// ░░▓▓░░░░░░░░░░░░░░▓▓░░
	// ░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░
	// ░░░░░░░░░░░░░░░░░░░░░░
}

func Example_gfx_RectFill() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 9, 11), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Translate(4, 5)

	gc.Color(1)
	gc.RectFill(-2, -4, 2, 4)

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
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░░░░░░░░░░░░░░░
}

func Example_gfx_RoundRect() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 17, 11), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Translate(8, 5)

	gc.Color(1)
	gc.RoundRect(-6, -4, 6, 4, 2)

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
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓░░░░░░░░░░░░░░░░░░▓▓░░░░░░
	// ░░░░▓▓░░░░░░░░░░░░░░░░░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░░░░░░░░░░░░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░░░░░░░░░░░░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░░░░░░░░░░░░░░░░░▓▓░░░░
	// ░░░░▓▓░░░░░░░░░░░░░░░░░░░░░░▓▓░░░░
	// ░░░░░░▓▓░░░░░░░░░░░░░░░░░░▓▓░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
}

func Example_gfx_RoundRectFill() {
	pal := color.Palette{
		color.NRGBA{0x00, 0x00, 0x00, 0xFF},
		color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	dst := image.NewPaletted(image.Rect(0, 0, 17, 11), pal)

	gc := canvas.New(dst)
	gc.Cls(0)

	gc.Translate(8, 5)

	gc.Color(1)
	gc.RoundRectFill(-6, -4, 6, 4, 2)

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
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
}
