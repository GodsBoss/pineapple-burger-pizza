package game

import (
	"strconv"

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

			offsetX := (160 - w*pizzaFieldWidth/2) * scale
			offsetY := (100 - h*pizzaFieldHeight/2) * scale

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
						offsetX+x*pizzaFieldWidth*scale,
						offsetY+y*pizzaFieldHeight*scale,
						scale,
						0,
					).Render(output)
				}
			}
		}

		// Render laying ingredients
		for _, ingredient := range d.waitingIngredients {
			var key canvas2drendering.SpriteKey
			var amountOffsetX int
			var amountOffsetY int

			switch ingredient.typ {
			case ingredientAnchovi:
				key = keys.ingredientAnchovi[0]
				amountOffsetX = 20
				amountOffsetY = 20
			}

			// Skip unknown ingredient types.
			if key == nil {
				continue
			}

			spriteMap.CreateSprite(
				key,
				canvas2drendering.SpriteAttributes{},
				ingredient.x*scale,
				ingredient.y*scale,
				scale,
				0,
			).Render(output)

			amountString := "*" + strconv.Itoa(ingredient.amount)

			tm.Create(
				(ingredient.x+amountOffsetX)*scale,
				(ingredient.y+amountOffsetY)*scale,
				scale,
				[]string{amountString},
			).Render(output)
		}
	}
}

const (
	pizzaFieldWidth  = 20
	pizzaFieldHeight = 16
)
