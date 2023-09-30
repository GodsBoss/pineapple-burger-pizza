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

		if pizzaKey, ok := keys.pizzas[d.pizza.Width()]; ok {
			spriteMap.CreateSprite(
				pizzaKey,
				canvas2drendering.SpriteAttributes{},
				centerOffsetX,
				centerOffsetY,
				scale,
				0,
			).Render(output)
		}

		for _, placed := range d.placedIngredients {
			offsetX := -ingredientSizes[placed.typ].Width / 2
			offsetY := -ingredientSizes[placed.typ].Height / 2

			spriteMap.CreateSprite(
				keys.ingredients[placed.typ][int(placed.orientation)],
				canvas2drendering.SpriteAttributes{},
				(placed.x+offsetX)*scale,
				(placed.y+offsetY)*scale,
				scale,
				0,
			).Render(output)
		}

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
			}
		}

		// Render laying ingredients
		for _, ingredient := range d.waitingIngredients {
			key := keys.ingredients[ingredient.typ][0]
			size := ingredientSizes[ingredient.typ]

			amountOffsetX := size.Width/2 - 5
			amountOffsetY := size.Height

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
				keys.ingredients[d.draggedIngredient.typ][int(d.draggedIngredient.orientation)],
				canvas2drendering.SpriteAttributes{},
				(d.draggedIngredient.x+offsetX)*scale,
				(d.draggedIngredient.y+offsetY)*scale,
				scale,
				0,
			).Render(output)

			for _, field := range d.draggedIngredient.validFields {
				spriteMap.CreateSprite(
					keys.ingredientGridOverlayFree,
					canvas2drendering.SpriteAttributes{},
					centerOffsetX+field.X()*pizzaFieldWidth*scale,
					centerOffsetY+field.Y()*pizzaFieldHeight*scale,
					scale,
					0,
				).Render(output)
			}

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

		// Render customer likes.
		pos := 0
		for _, fl := range flavorList {
			if amount, ok := d.customer.likes[fl]; ok {
				spriteMap.CreateSprite(
					keys.flavors[fl],
					canvas2drendering.SpriteAttributes{},
					20*scale,
					(20+pos*18)*scale,
					scale,
					0,
				).Render(output)
				tm.Create(
					40*scale,
					(25+pos*18)*scale,
					scale,
					[]string{"*" + strconv.Itoa(amount)},
				).Render(output)
				pos++
			}
		}
		if len(d.customer.likes) > 0 {
			spriteMap.CreateSprite(
				keys.customerLike,
				canvas2drendering.SpriteAttributes{},
				20*scale,
				2*scale,
				scale,
				0,
			).Render(output)
		}

		// Render customer dislikes.
		pos = 0
		for _, fl := range flavorList {
			if _, ok := d.customer.dislikes[fl]; ok {
				spriteMap.CreateSprite(
					keys.flavors[fl],
					canvas2drendering.SpriteAttributes{},
					(140+24*pos)*scale,
					30*scale,
					scale,
					0,
				).Render(output)
				pos++
			}
		}
		if len(d.customer.dislikes) > 0 {
			spriteMap.CreateSprite(
				keys.customerDislike,
				canvas2drendering.SpriteAttributes{},
				116*scale,
				30*scale,
				scale,
				0,
			).Render(output)
		}

		// Render pizza flavors.
		pos = 0
		for _, fl := range flavorList {
			if amount, ok := d.pizza.flavors[fl]; ok {
				spriteMap.CreateSprite(
					keys.flavors[fl],
					canvas2drendering.SpriteAttributes{},
					(160+w*pizzaFieldWidth/2)*scale,
					(100-h*pizzaFieldHeight/2+pos*18)*scale,
					scale,
					0,
				).Render(output)
				tm.Create(
					(180+w*pizzaFieldWidth/2)*scale,
					(105-h*pizzaFieldHeight/2+pos*18)*scale,
					scale,
					[]string{"*" + strconv.Itoa(amount)},
				).Render(output)
				pos++
			}
		}

		// Render reputation
		for y := 0; y < 10; y++ {
			key := keys.reputationOK
			if d.reputation <= y {
				key = keys.reputationGone
			}
			spriteMap.CreateSprite(
				key,
				canvas2drendering.SpriteAttributes{},
				280*scale,
				(10+18*y)*scale,
				scale,
				0,
			).Render(output)
		}
	}
}

const (
	pizzaFieldWidth  = 20
	pizzaFieldHeight = 16
)
