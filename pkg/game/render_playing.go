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

		renderSprite(keys.backgroundPlaying, 0, 0, 0)

		w := d.pizza.Width()
		h := d.pizza.Height()

		centerOffsetX := (160 - w*pizzaFieldWidth/2)
		centerOffsetY := (100 - h*pizzaFieldHeight/2)

		if pizzaKey, ok := keys.pizzas[d.pizza.Width()]; ok {
			renderSprite(pizzaKey, centerOffsetX, centerOffsetY, 0)
		}

		for _, placed := range d.placedIngredients {
			offsetX := -placed.width / 2
			offsetY := -placed.height / 2

			renderSprite(keys.ingredients[placed.typ][int(placed.orientation)], placed.x+offsetX, placed.y+offsetY, 0)
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

					renderSprite(overlayKey, centerOffsetX+x*pizzaFieldWidth, centerOffsetY+y*pizzaFieldHeight, 0)
				}
			}
		}

		// Render customer
		customerHeadKey := keys.customerHeadNormal
		switch d.customer.state {
		case customerStateAngry:
			customerHeadKey = keys.customerHeadAngry
		case customerStateHappy:
			customerHeadKey = keys.customerHeadHappy
		}
		renderSprite(customerHeadKey, 35, 0, 0)

		// Render laying ingredients
		for _, ingredient := range d.waitingIngredients {
			key := keys.ingredients[ingredient.typ][0]
			size := ingredientSizes[ingredient.typ]

			amountOffsetX := size.Width/2 - 5
			amountOffsetY := size.Height

			renderSprite(key, ingredient.x, ingredient.y, 0)

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

		// Render customer likes.
		pos := 0
		for _, fl := range flavorList {
			if amount, ok := d.customer.likes[fl]; ok {
				renderSprite(keys.flavors[fl], 3, 20+pos*18, 0)
				tm.Create(
					20*scale,
					(25+pos*18)*scale,
					scale,
					[]string{"*" + strconv.Itoa(amount)},
				).Render(output)
				pos++
			}
		}
		if len(d.customer.likes) > 0 {
			renderSprite(keys.customerLike, 3, 2, 0)
		}

		// Render customer dislikes.
		pos = 0
		for _, fl := range flavorList {
			if _, ok := d.customer.dislikes[fl]; ok {
				renderSprite(keys.flavors[fl], 92, 20+pos*18, 0)
				pos++
			}
		}
		if len(d.customer.dislikes) > 0 {
			renderSprite(keys.customerDislike, 92, 2, 0)
		}

		// Render pizza flavors.
		pos = 0
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
