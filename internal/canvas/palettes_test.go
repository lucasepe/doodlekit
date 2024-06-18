package canvas

import (
	"image/color"
	"testing"
)

func TestPalettes(t *testing.T) {
	for _, tc := range []struct {
		Name       string
		Palette    color.Palette
		ColorCount int
	}{

		{"1Bit", Palette1Bit, 2},

		{"2BitGrayScale", Palette2BitGrayScale, 4},
		{"3Bit", Palette3Bit, 8},
		{"CGA", PaletteCGA, 16},
		{"Famicube", PaletteFamicube, 64},
		{"Ink", PaletteInk, 5},
		{"PICO8", PalettePICO8, 16},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			if got, want := len(tc.Palette), tc.ColorCount; got != want {
				t.Fatalf("unexpected number of colors: %d, want %d", got, want)
			}

			for n, want := range tc.Palette {
				if got := Color(tc.Palette, n); got != want {
					t.Fatalf("Color(%d) = %v, want %v", n, got, want)
				}
			}
		})
	}
}

func TestPalettesByNameAndCount(t *testing.T) {
	var nc int

	for _, p := range PalettesByNumberOfColors {
		nc += len(p)
	}

	pc := len(PaletteByName)

	if nc != pc {
		t.Fatalf("nc = %d, want %d", nc, pc)
	}
}
