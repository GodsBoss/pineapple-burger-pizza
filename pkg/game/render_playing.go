package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
)

func renderPlaying(spriteMap canvas2drendering.SpriteMap, keys spriteKeys, tm *textManager) stateRendererFunc {
	return func(output *dom.Context2D, d *data, scale int) {
		spriteMap.CreateSprite(
			keys.backgroundPlaying,
			canvas2drendering.SpriteAttributes{},
			0,
			0,
			scale,
			0,
		).Render(output)

		// Render pizza grid
		if d.pizzaGridOverlayVisible {
			w := d.pizza.Width()
			h := d.pizza.Height()

			offsetX := (160 - w*8) * scale
			offsetY := (100 - h*8) * scale

			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					if d.pizza.grid[x][y].invalid {
						continue
					}

					overlayKey := keys.pizzaGridOverlayFree
					if d.pizza.grid[x][y].occupied {
						overlayKey = keys.pizzaGridOverlayOccupied
					}

					spriteMap.CreateSprite(
						overlayKey,
						canvas2drendering.SpriteAttributes{},
						offsetX+x*16*scale,
						offsetY+y*16*scale,
						scale,
						0,
					).Render(output)
				}
			}
		}
	}
}
