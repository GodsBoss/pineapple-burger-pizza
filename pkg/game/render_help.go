package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
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
	}
}
