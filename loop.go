package doodlekit

import (
	"context"
	"fmt"
	"path/filepath"
	"time"
)

func NewLoop(id string, opts ...Option) Loop {
	res := &demoLoop{
		id:        id,
		fps:       25.0,
		stopAfter: 6,
		x2:        false,
	}

	for _, o := range opts {
		o(res)
	}

	return res
}

type Option func(l *demoLoop)

func FPS(fps float64) Option {
	return func(l *demoLoop) {
		l.fps = fps
	}
}

func StopAfter(secs int) Option {
	return func(l *demoLoop) {
		l.stopAfter = secs
	}
}

func OutDir(dir string) Option {
	return func(l *demoLoop) {
		l.outdir = dir
	}
}

func X2() Option {
	return func(l *demoLoop) {
		l.x2 = true
	}
}

type Scene interface {
	Init(ctx context.Context)
	Update(ctx context.Context, dt float64)
	Draw(ctx context.Context)
}

type Loop interface {
	Run(fx []Scene)
}

var _ Loop = (*demoLoop)(nil)

type demoLoop struct {
	id        string
	fps       float64
	stopAfter int
	outdir    string
	x2        bool
}

func (dl *demoLoop) Run(fx []Scene) {
	opts := []option{}
	if dl.x2 {
		opts = append(opts, x2())
	}
	ctx := newContext(opts...)

	for _, ds := range fx {
		ds.Init(ctx)
	}

	ticker := time.NewTicker(time.Second / time.Duration(dl.fps))
	defer ticker.Stop()

	exitTimer := time.NewTimer(time.Duration(dl.stopAfter) * time.Second)
	defer exitTimer.Stop()

	for {
		select {
		case <-ticker.C:
			dt := 1.0 / dl.fps

			for _, ds := range fx {
				ds.Update(ctx, dt)
				ds.Draw(ctx)
			}

			if gc := Canvas(ctx); gc != nil {
				gc.Record()
			}

		case <-exitTimer.C:
			gc := Canvas(ctx)
			if gc == nil {
				return
			}

			if len(dl.id) == 0 {
				dl.id = time.Now().Format("200601021504")
			}

			if len(dl.outdir) > 0 {
				gc.Save(filepath.Join(dl.outdir, fmt.Sprintf("%s.gif", dl.id)))
			} else {
				gc.Save(fmt.Sprintf("%s.gif", dl.id))
			}
			return
		}
	}
}
