package canvas_test

import (
	"fmt"
	"os"

	"github.com/lucasepe/doodlekit/internal/canvas"
	"github.com/lucasepe/doodlekit/internal/fonts"
)

func Example_canvas_MustLoadFont() {
	fnt := canvas.MustLoadFont(fonts.Micro())

	fmt.Printf("Name: %s\n", fnt.Name)
	fmt.Printf("Size: %d\n", fnt.Size)

	// Output:
	//
	// Name: micro
	// Size: -1
}

func Example_canvas_GlyphBounds() {
	r := rune('A')

	face := canvas.MustLoadFont(fonts.Micro()).NewFace()

	bounds, advance, ok := face.GlyphBounds(r)
	if !ok {
		fmt.Fprintf(os.Stderr, "the face does not contain a glyph for '%v'", r)
		return
	}

	w := canvas.UnfixI(bounds.Max.X) - canvas.UnfixI(bounds.Min.X)
	h := canvas.UnfixI(bounds.Max.Y) - canvas.UnfixI(bounds.Min.Y)
	fmt.Printf("GlyphBounds: %dx%d\n", w, h)
	fmt.Printf("Advance: %v\n", canvas.UnfixI(advance))

	// Output:
	//
	// 	GlyphBounds: 4x5
	// Advance: 4
}
