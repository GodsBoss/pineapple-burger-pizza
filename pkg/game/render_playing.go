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

		w := d.pizza.Width()
		h := d.pizza.Height()

		centerOffsetX := (160 - w*pizzaFieldWidth/2) * scale
		centerOffsetY := (100 - h*pizzaFieldHeight/2) * scale

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if d.pizza.grid[x][y].invalid {
					continue
				}

				if d.pizzaGridOverlayVisible {
					overlayKey := keys.pizzaGridOverlayFree
					if d.pizza.grid[x][y].occupied {
						overlayKey = keys.pizzaGridOverlayOccupied
					}

					spriteMap.CreateSprite(
						overlayKey,
						canvas2drendering.SpriteAttributes{},
						centerOffsetX+x*pizzaFieldWidth*scale,
						centerOffsetY+y*pizzaFieldHeight*scale,
						scale,
						0,
					).Render(output)
				}

				if d.pizza.grid[x][y].draggedIngredientTarget {
					overlayKey := keys.ingredientGridOverlayFree
					if d.pizza.grid[x][y].occupied {
						overlayKey = keys.ingredientGridOverlayOccupied
					}

					spriteMap.CreateSprite(
						overlayKey,
						canvas2drendering.SpriteAttributes{},
						centerOffsetX+x*pizzaFieldWidth*scale,
						centerOffsetY+y*pizzaFieldHeight*scale,
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

		// Render dragged ingredient
		if d.draggedIngredient != nil {
			offsetX := -d.draggedIngredient.Width() / 2
			offsetY := -d.draggedIngredient.Height() / 2

			spriteMap.CreateSprite(
				keys.ingredientAnchovi[int(d.draggedIngredient.orientation)],
				canvas2drendering.SpriteAttributes{},
				(d.draggedIngredient.x+offsetX)*scale,
				(d.draggedIngredient.y+offsetY)*scale,
				scale,
				0,
			).Render(output)

			for _, field := range d.draggedIngredient.invalidFields {
				spriteMap.CreateSprite(
					keys.ingredientGridOverlayOccupied,
					canvas2drendering.SpriteAttributes{},
					centerOffsetX+field.X()*pizzaFieldWidth*scale,
					centerOffsetY+field.Y()*pizzaFieldHeight*scale,
					scale,
					0,
				).Render(output)
			}
		}
	}
}

const (
	pizzaFieldWidth  = 20
	pizzaFieldHeight = 16
)
