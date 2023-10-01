package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

const helpState = "help"

func initHelp(d *data) game.NextState {
	d.state = helpState
	d.pizza = createPizza(5)
	d.reputation = 5
	d.waitingIngredients = []waitingIngredient{
		{
			typ:    ingredientAnchovi,
			amount: 3,
			x:      220 - ingredientSizes[ingredientAnchovi].Width/2,
			y:      180 - ingredientSizes[ingredientAnchovi].Height,
		},
	}

	return game.SameState()
}

func createReceiveKeyEventHelp(title game.StateID, playing game.StateID) func(d *data, event keyboard.Event) game.NextState {
	return func(d *data, event keyboard.Event) game.NextState {
		if event.Key == "t" || event.Key == "T" {
			return game.SwitchState(title)
		}

		if event.Key == "p" || event.Key == "P" {
			return game.SwitchState(playing)
		}

		return game.SameState()
	}
}
