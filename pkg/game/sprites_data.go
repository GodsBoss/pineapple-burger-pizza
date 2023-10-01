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
	"background_table": {
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
	"ingredient_grid_overlay_free": {
		X:      0,
		Y:      48,
		Width:  20,
		Height: 16,
	},
	"ingredient_grid_overlay_occupied": {
		X:      0,
		Y:      32,
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
	"ananas_0": {
		X:      20,
		Y:      32,
		Width:  60,
		Height: 48,
	},
	"ananas_1": {
		X:      80,
		Y:      32,
		Width:  60,
		Height: 48,
	},
	"ananas_2": {
		X:      140,
		Y:      32,
		Width:  60,
		Height: 48,
	},
	"ananas_3": {
		X:      200,
		Y:      32,
		Width:  60,
		Height: 48,
	},
	"rubber_boots_0": {
		X:      0,
		Y:      80,
		Width:  40,
		Height: 32,
	},
	"rubber_boots_1": {
		X:      40,
		Y:      80,
		Width:  40,
		Height: 32,
	},
	"rubber_boots_2": {
		X:      80,
		Y:      80,
		Width:  40,
		Height: 32,
	},
	"rubber_boots_3": {
		X:      120,
		Y:      80,
		Width:  40,
		Height: 32,
	},
	"burger_0": {
		X:      0,
		Y:      160,
		Width:  60,
		Height: 32,
	},
	"burger_1": {
		X:      0,
		Y:      192,
		Width:  60,
		Height: 32,
	},
	"burger_2": {
		X:      60,
		Y:      176,
		Width:  40,
		Height: 48,
	},
	"burger_3": {
		X:      100,
		Y:      176,
		Width:  40,
		Height: 48,
	},
	"customer_like": {
		X:      400,
		Y:      0,
		Width:  16,
		Height: 16,
	},
	"customer_dislike": {
		X:      400,
		Y:      16,
		Width:  16,
		Height: 16,
	},
	"flavor_sweet": {
		X:      416,
		Y:      0,
		Width:  16,
		Height: 16,
	},
	"flavor_calamari": {
		X:      432,
		Y:      0,
		Width:  16,
		Height: 16,
	},
	"flavor_salty": {
		X:      448,
		Y:      0,
		Width:  16,
		Height: 16,
	},
	"flavor_fish": {
		X:      464,
		Y:      0,
		Width:  16,
		Height: 16,
	},
	"reputation_ok": {
		X:      0,
		Y:      1020,
		Width:  16,
		Height: 16,
	},
	"reputation_gone": {
		X:      16,
		Y:      1020,
		Width:  16,
		Height: 16,
	},
	"pizza_3": {
		X:      0,
		Y:      112,
		Width:  60,
		Height: 48,
	},
	"pizza_4": {
		X:      60,
		Y:      112,
		Width:  80,
		Height: 64,
	},
	"pizza_5": {
		X:      140,
		Y:      112,
		Width:  100,
		Height: 80,
	},
	"customer_head_normal": {
		X:      1520,
		Y:      200,
		Width:  40,
		Height: 32,
	},
	"customer_head_angry": {
		X:      1560,
		Y:      200,
		Width:  40,
		Height: 32,
	},
	"customer_head_happy": {
		X:      1520,
		Y:      232,
		Width:  40,
		Height: 32,
	},
	"customer_body": {
		X:      1517,
		Y:      266,
		Width:  82,
		Height: 61,
	},
	"help_icon": {
		X:      400,
		Y:      32,
		Width:  16,
		Height: 16,
	},
	"help_icon_active": {
		X:      400,
		Y:      48,
		Width:  16,
		Height: 16,
	},
}
