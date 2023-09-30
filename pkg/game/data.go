package game

// data represents all of the game's data.
type data struct {
	// state is the state the game is in (title, playing, etc.). Set by states, used for rendering.
	state string

	// pizza is the pizza visible when playing.
	pizza *pizza

	// pizzaGridOverlayVisible determines whether the pizza grid overlay is visible.
	pizzaGridOverlayVisible bool
}
