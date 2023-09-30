package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
	r "github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
)

func createSpriteMap(sourceImage *dom.Image) (canvas2drendering.SpriteMap, spriteKeys) {
	spriteMap := canvas2drendering.NewSpriteMap(sourceImage)
	spr := spriteKeys{}
	return spriteMap, spr
}

// spriteKeys holds the sprite keys generated when adding sprite specs to the spriteKeys factory.
type spriteKeys struct {
	backgroundTitle    r.SpriteKey
	backgroundPlaying  r.SpriteKey
	backgroundGameOver r.SpriteKey
}
