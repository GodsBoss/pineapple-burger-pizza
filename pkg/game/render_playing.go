package game

import (
	"strconv"

	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
	"github.com/GodsBoss/pineapple-burger-pizza/pkg/minifmt"
)

func renderPlaying(spriteMap canvas2drendering.SpriteMap, keys spriteKeys, tm *textManager) stateRendererFunc {
	return func(output *dom.Context2D, d *data, scale int) {
		renderSprite := createRenderSprite(spriteMap, output, scale)
		renderText := createRenderText(tm, output, scale)

		renderSprite(keys.backgroundPlaying, 0, 0, 0)

		w := d.pizza.Width()
		h := d.pizza.Height()

		centerOffsetX := (160 - w*pizzaFieldWidth/2)
		centerOffsetY := (100 - h*pizzaFieldHeight/2)

		renderPizza(renderSprite, keys, *d.pizza)
		renderPlacedIngredients(renderSprite, keys, d.placedIngredients)

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

					renderSprite(overlayKey, centerOffsetX+x*pizzaFieldWidth, centerOffsetY+y*pizzaFieldHeight, 0)
				}
			}
		}

		// Render customer
		customerHeadKey := keys.customerHeadNormal
		switch d.customer.mood {
		case customerMoodAngry:
			customerHeadKey = keys.customerHeadAngry
		case customerMoodHappy:
			customerHeadKey = keys.customerHeadHappy
		}
		renderSprite(keys.customerBody, 14, 22, 0)
		renderSprite(customerHeadKey, 35, 0, 0)

		renderWaitingIngredients(renderSprite, keys, renderText, d.waitingIngredients)

		// Render dragged ingredient
		if d.draggedIngredient != nil {
			offsetX := -d.draggedIngredient.Width() / 2
			offsetY := -d.draggedIngredient.Height() / 2

			renderSprite(
				keys.ingredients[d.draggedIngredient.typ][int(d.draggedIngredient.orientation)],
				d.draggedIngredient.x+offsetX,
				d.draggedIngredient.y+offsetY,
				0,
			)

			for _, field := range d.draggedIngredient.validFields {
				renderSprite(
					keys.ingredientGridOverlayFree,
					centerOffsetX+field.X()*pizzaFieldWidth,
					centerOffsetY+field.Y()*pizzaFieldHeight,
					0,
				)
			}

			for _, field := range d.draggedIngredient.invalidFields {
				renderSprite(
					keys.ingredientGridOverlayOccupied,
					centerOffsetX+field.X()*pizzaFieldWidth,
					centerOffsetY+field.Y()*pizzaFieldHeight,
					0,
				)
			}
		}

		renderCustomerLikes(renderSprite, keys, renderText, d.customer.likes)
		renderCustomerDislikes(renderSprite, keys, d.customer.dislikes)

		// Render pizza flavors.
		pos := 0
		for _, fl := range flavorList {
			pizzaAmount, isOnPizza := d.pizza.flavors[fl]
			var ingredientAmount int
			if d.draggedIngredient != nil {
				ingredientAmount = ingredientFlavors[d.draggedIngredient.typ][fl]
			}
			if isOnPizza || ingredientAmount > 0 {
				renderSprite(keys.flavors[fl], 160+w*pizzaFieldWidth/2, 100-h*pizzaFieldHeight/2+pos*18, 0)
				content := ""
				if pizzaAmount > 0 {
					content += "*" + strconv.Itoa(pizzaAmount)
				}
				if ingredientAmount > 0 {
					content += "+" + strconv.Itoa(ingredientAmount)
				}
				tm.Create(
					(180+w*pizzaFieldWidth/2)*scale,
					(105-h*pizzaFieldHeight/2+pos*18)*scale,
					scale,
					[]string{content},
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
			renderSprite(key, 300, 10+18*y, 0)
		}

		// Render score
		tm.Create(200*scale, 2*scale, scale, []string{"score: " + minifmt.FormatInt(d.score, 6, " ")}).Render(output)

		// Render hiscore
		if d.highscore > 0 {
			tm.Create(188*scale, 10*scale, scale, []string{"hiscore: " + minifmt.FormatInt(d.highscore, 6, " ")}).Render(output)
		}
	}
}

const (
	pizzaFieldWidth  = 20
	pizzaFieldHeight = 16
)
