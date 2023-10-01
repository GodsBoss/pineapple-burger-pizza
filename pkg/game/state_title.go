package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

const titleState = "title"

func initTitle(d *data) game.NextState {
	d.state = titleState

	return game.SameState()
}

func createReceiveTitle(helpState game.StateID, playingState game.StateID) func(d *data, ev keyboard.Event) game.NextState {
	return func(d *data, ev keyboard.Event) game.NextState {
		if ev.Key == "h" || ev.Key == "H" {
			return game.SwitchState(helpState)
		}

		if ev.Key == "p" {
			return game.SwitchState(playingState)
		}

		return game.SameState()
	}
}

func createReceiveTickEventTitle() func(d *data, event tick.Event) game.NextState {
	return func(d *data, event tick.Event) game.NextState {
		return game.SameState()
	}
}
