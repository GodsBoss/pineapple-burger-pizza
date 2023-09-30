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
	"pizza_grid_overlay_free": {
		X:      0,
		Y:      16,
		Width:  20,
		Height: 16,
	},
	"pizza_grid_overlay_occupied": {
		X:      0,
		Y:      0,
		Width:  20,
		Height: 16,
	},
	"anchovy_0": {
		X:      20,
		Y:      0,
		Width:  40,
		Height: 16,
	},
	"anchovy_1": {
		X:      60,
		Y:      0,
		Width:  20,
		Height: 32,
	},
	"anchovy_2": {
		X:      20,
		Y:      16,
		Width:  40,
		Height: 16,
	},
	"anchovy_3": {
		X:      80,
		Y:      0,
		Width:  20,
		Height: 32,
	},
}
