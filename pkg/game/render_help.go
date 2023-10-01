package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
	"github.com/GodsBoss/pineapple-burger-pizza/pkg/text"
)

func renderHelp(spriteMap canvas2drendering.SpriteMap, keys spriteKeys, tm *textManager) stateRendererFunc {
	return func(output *dom.Context2D, d *data, scale int) {
		renderSprite := createRenderSprite(spriteMap, output, scale)
		renderText := createRenderText(tm, output, scale)

		renderSprite(keys.backgroundTable, 0, 0, 0)

		renderSprite(keys.customerBody, 14, 22, 0)
		renderSprite(keys.customerHeadNormal, 35, 0, 0)
		renderPizza(renderSprite, keys, *d.pizza)
		renderWaitingIngredients(renderSprite, keys, renderText, d.waitingIngredients)
		renderReputation(renderSprite, keys, d.reputation)
		renderCustomerLikes(renderSprite, keys, renderText, d.customer.likes)
		renderCustomerDislikes(renderSprite, keys, d.customer.dislikes)

		for _, button := range d.helpButtons {
			key := keys.helpIcon
			if button.active {
				key = keys.helpIconActive
			}
			renderSprite(key, button.x, button.y, 0)
		}

		if d.helpText != "" {
			renderText(1, 90, text.Lines(18, d.helpText))
		}

		renderSprite(keys.keyboard, 150, 18, 0)
	}
}
