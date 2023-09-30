package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

const gameOverState = "game_over"

func initGameOver(d *data) game.NextState {
	d.state = gameOverState

	return game.SameState()
}

func createReceiveKeyEventGameOver(titleState game.StateID) func(d *data, event keyboard.Event) game.NextState {
	return func(d *data, event keyboard.Event) game.NextState {
		if event.Key == "t" || event.Key == "T" {
			return game.SwitchState(titleState)
		}
		return game.SameState()
	}
}
