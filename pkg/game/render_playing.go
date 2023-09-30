package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
)

func renderPlaying(spriteMap canvas2drendering.SpriteMap, keys spriteKeys) stateRendererFunc {
	return func(output *dom.Context2D, d *data, scale int) {
		s := spriteMap.CreateSprite(
			keys.backgroundPlaying,
			canvas2drendering.SpriteAttributes{},
			0,
			0,
			scale,
			0,
		)
		s.Render(output)
	}
}
