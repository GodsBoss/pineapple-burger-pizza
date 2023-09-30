package game

import (
	"strings"

	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
)

func newTextManager(spriteManager canvas2drendering.SpriteMap, charSprites map[byte]canvas2drendering.SpriteData) *textManager {
	charKeyMap := make(map[byte]canvas2drendering.SpriteKey)

	for ch, spriteData := range charSprites {
		charKeyMap[ch] = spriteManager.AddSpriteSpecification(
			canvas2drendering.SpriteSpecification{
				canvas2drendering.SpriteAttributes{}: {
					X:      spriteData.X,
					Y:      spriteData.Y,
					Width:  spriteData.Width,
					Height: spriteData.Height,
				},
			},
		)
	}

	return &textManager{
		spriteManager: spriteManager,
		charKeyMap:    charKeyMap,
	}
}

type textManager struct {
	spriteManager canvas2drendering.SpriteMap
	charKeyMap    map[byte]canvas2drendering.SpriteKey
}

func (tm *textManager) Create(x int, y int, scale int, lines []string) canvas2drendering.Renderable {
	return renderableText{
		manager: tm,
		x:       x,
		y:       y,
		scale:   scale,
		lines:   lines,
	}
}

type renderableText struct {
	manager *textManager
	x       int
	y       int
	scale   int
	lines   []string
}

func (t renderableText) Render(output *dom.Context2D) {
	for i := range t.lines {
		for p := range t.lines[i] {
			key, ok := t.manager.charKeyMap[toLowerChar(t.lines[i][p])]
			if !ok {
				continue
			}

			x := t.x + p*(charWidth+1)*t.scale
			y := t.y + i*(charHeight+1)*t.scale

			t.manager.spriteManager.CreateSprite(key, canvas2drendering.SpriteAttributes{}, x, y, t.scale, 0).Render(output)
		}
	}
}

const (
	charWidth  = 5
	charHeight = 7
)

func toLowerChar(b byte) byte {
	lower := strings.ToLower(string(b))
	return lower[0]
}
