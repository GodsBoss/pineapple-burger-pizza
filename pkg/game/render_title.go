package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
)

func renderTitle(spriteMap canvas2drendering.SpriteMap, keys spriteKeys, tm *textManager) stateRendererFunc {
	return func(output *dom.Context2D, d *data, scale int) {
		renderSprite := createRenderSprite(spriteMap, output, scale)
		renderText := createRenderText(tm, output, scale)

		s := spriteMap.CreateSprite(
			keys.backgroundTitle,
			canvas2drendering.SpriteAttributes{},
			0,
			0,
			scale,
			0,
		)
		s.Render(output)

		renderText(5, 188, []string{"Press 'H' for help"})
		if d.readyToPlay {
			renderText(5, 177, []string{"Press 'P' to play"})
		}

		currentIngredient := ingredientList[d.titleIngredientIndex]
		renderSprite(
			keys.ingredients[currentIngredient][0],
			160-ingredientSizes[currentIngredient].Width/2,
			115-ingredientSizes[currentIngredient].Height/2,
			0,
		)
	}
}
