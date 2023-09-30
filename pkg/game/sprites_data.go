package game

import "github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"

// spritesData contains a map of all the sprites this game uses.
var spritesData = map[string]canvas2drendering.SpriteData{
	"background_title": {
		X:      1600,
		Y:      0,
		Width:  320,
		Height: 200,
	},
	"background_playing": {
		X:      1600,
		Y:      200,
		Width:  320,
		Height: 200,
	},
	"background_game_over": {
		X:      1600,
		Y:      400,
		Width:  320,
		Height: 200,
	},
}
