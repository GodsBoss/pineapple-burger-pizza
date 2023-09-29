package game

import "github.com/GodsBoss/gggg/v2/pkg/game"

const gameOverState = "game_over"

func initGameOver(d *data) game.NextState {
	d.state = gameOverState

	return game.SameState()
}
