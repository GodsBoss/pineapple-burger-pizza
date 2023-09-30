package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

const playingState = "playing"

func initPlaying(d *data) game.NextState {
	d.state = playingState
	d.pizzaGridOverlayVisible = false
	d.pizza = createPizza(5)

	return game.SameState()
}

func createReceiveKeyEventPlaying() func(d *data, event keyboard.Event) game.NextState {
	return func(d *data, event keyboard.Event) game.NextState {
		if event.Key == "g" && keyboard.IsUpEvent(event) {
			d.pizzaGridOverlayVisible = !d.pizzaGridOverlayVisible
		}

		return game.SameState()
	}
}

// createPizza creates a pizza which is basically a diameter x diameter grid, but the corner pieces are missing.
// For small diameters, this is "good enough".
func createPizza(diameter int) *pizza {
	grid := make([][]pizzaField, diameter)

	for i := range grid {
		grid[i] = make([]pizzaField, diameter)
	}

	max := diameter - 1

	// Disable corner fields.
	grid[0][0].invalid = true
	grid[max][0].invalid = true
	grid[0][max].invalid = true
	grid[max][max].invalid = true

	return &pizza{
		grid: grid,
	}
}

type pizza struct {
	grid [][]pizzaField
}

func (p pizza) Width() int {
	return len(p.grid)
}

func (p pizza) Height() int {
	return len(p.grid[0])
}

type pizzaField struct {
	// invalid marks a field as invalid for placing an ingredient.
	invalid bool

	// occupied marks whether part of an ingredient occupies this field.
	occupied bool
}
