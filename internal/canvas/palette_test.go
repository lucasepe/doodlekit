package canvas

import (
	"fmt"
	"image/color"
	"testing"
)

func TestPalette(t *testing.T) {
	pal := doodleKitPalette()

	if got := len(pal); got != 17 {
		t.Fatalf("unexpected number of colors: %d, want 17", got)
	}

	cases := []struct {
		idx  int
		want color.Color
	}{
		{0, color.NRGBA{0xFF, 0xFF, 0xFF, 0x00}},
		{3, color.NRGBA{0x7E, 0x25, 0x53, 0xFF}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc[%d]", i), func(t *testing.T) {
			if got := Color(pal, tc.idx); got != tc.want {
				t.Fatalf("Color(%d) = %v, want %v", tc.idx, got, tc.want)
			}
		})
	}
}
