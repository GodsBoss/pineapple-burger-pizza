package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
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
	r.output.DisableImageSmoothing()
	w, h := r.output.Size()
	r.output.ClearRect(0, 0, w, h)
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

type renderSpriteFunc func(key canvas2drendering.SpriteKey, x int, y int, frame int)

func createRenderSprite(spriteMap canvas2drendering.SpriteMap, output *dom.Context2D, scale int) func(key canvas2drendering.SpriteKey, x int, y int, frame int) {
	return func(key canvas2drendering.SpriteKey, x int, y int, frame int) {
		spriteMap.CreateSprite(key, canvas2drendering.SpriteAttributes{}, x*scale, y*scale, scale, frame).Render(output)
	}
}

func renderPizza(renderSprite renderSpriteFunc, keys spriteKeys, p pizza) {
	if pizzaKey, ok := keys.pizzas[p.Width()]; ok {
		renderSprite(
			pizzaKey,
			(160 - p.Width()*pizzaFieldWidth/2),
			(100 - p.Height()*pizzaFieldHeight/2),
			0,
		)
	}
}

func renderPlacedIngredients(renderSprite renderSpriteFunc, keys spriteKeys, placeds []placedIngredient) {
	for _, placed := range placeds {
		offsetX := -placed.width / 2
		offsetY := -placed.height / 2

		renderSprite(keys.ingredients[placed.typ][int(placed.orientation)], placed.x+offsetX, placed.y+offsetY, 0)
	}
}
