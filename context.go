package brio

import (
	"context"

	"github.com/lucasepe/doodlekit/internal/canvas"
	"github.com/lucasepe/doodlekit/internal/rng"
)

// Canvas return the drawing context pointer from the context.
func Canvas(ctx context.Context) *canvas.Canvas {
	v := ctx.Value(contextKeyCanvas)
	if val, ok := v.(canvas.Canvas); ok {
		return &val
	}
	return nil
}

// Rng return the pseudo random numbers generator from the context.
func Rng(ctx context.Context) rng.RNG {
	v := ctx.Value(contextKeyRng)
	if val, ok := v.(rng.RNG); ok {
		return val
	}
	return nil
}

type contextKey string

func (c contextKey) String() string {
	return "doodlekit." + string(c)
}

var (
	contextKeyCanvas = contextKey("canvas")
	contextKeyRng    = contextKey("rng")
)

type option func(*canvas.Canvas)

func x2() option {
	return func(gc *canvas.Canvas) {
		gc.X2()
	}
}

func newContext(opts ...option) context.Context {
	gc := canvas.New(nil)
	for _, fn := range opts {
		fn(&gc)
	}

	ctx := context.WithValue(context.Background(), contextKeyCanvas, gc)
	ctx = context.WithValue(ctx, contextKeyRng, rng.New())

	return ctx
}
