package game

import "github.com/GodsBoss/gggg/v2/pkg/game"

const playingState = "playing"

func initPlaying(d *data) game.NextState {
	d.state = playingState

	return game.SameState()
}
