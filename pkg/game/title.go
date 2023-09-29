package game

import "github.com/GodsBoss/gggg/v2/pkg/game"

const titleState = "title"

func initTitle(d *data) game.NextState {
	d.state = titleState

	return game.SameState()
}
