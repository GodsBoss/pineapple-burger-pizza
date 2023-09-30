package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
)

type renderer struct {
	scaler scaler
	output *dom.Context2D

	stateRenderers map[string]stateRenderer
}

func (r *renderer) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	r.scaler.Recalculate(availableWidth, availableHeight)

	rw, rh := r.scaler.RealSize()
	s := float64(r.scaler.Scale())

	return rw, rh, s, s
}

func (r *renderer) SetOutput(ctx2d *dom.Context2D) {
	ctx2d.DisableImageSmoothing()
	r.output = ctx2d
}

func (r *renderer) Render(d *data) {
	if sr, ok := r.stateRenderers[d.state]; ok {
		sr.Render(r.output, d, r.scaler.Scale())
	}
}

func (r *renderer) AddStateRenderer(stateID string, sr stateRenderer) *renderer {
	if r.stateRenderers == nil {
		r.stateRenderers = make(map[string]stateRenderer)
	}
	r.stateRenderers[stateID] = sr
	return r
}

type stateRenderer interface {
	Render(output *dom.Context2D, d *data, scale int)
}

type stateRendererFunc func(output *dom.Context2D, d *data, scale int)

func (f stateRendererFunc) Render(output *dom.Context2D, d *data, scale int) {
	f(output, d, scale)
}
