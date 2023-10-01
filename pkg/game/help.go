package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

const helpState = "help"

func initHelp(d *data) game.NextState {
	d.state = helpState
	d.readyToPlay = true
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
	d.customer = &customer{
		likes: map[flavor]int{
			flavorSalty:    2,
			flavorCalamari: 1,
		},
		dislikes: map[flavor]struct{}{
			flavorSweet: struct{}{},
			flavorFish:  struct{}{},
		},
	}
	d.helpText = ""
	d.helpButtons = []helpButton{
		{
			id: "customer",
			x:  48,
			y:  32,
		},
		{
			id: "reputation",
			x:  280,
			y:  13,
		},
		{
			id: "ingredients",
			x:  210,
			y:  140,
		},
		{
			id: "hotkeys",
			x:  194,
			y:  17,
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

func createReceiveMouseEventHelp() func(d *data, event mouse.Event) game.NextState {
	return func(d *data, event mouse.Event) game.NextState {
		if mouse.IsPrimaryButtonEvent(event) && mouse.IsUpEvent(event) {
			for i, button := range d.helpButtons {
				if button.isInside(event.X, event.Y) {
					currentButtonState := button.active
					disableAllButtons(d.helpButtons)
					d.helpButtons[i].active = !currentButtonState
					d.helpText = ""
					if d.helpButtons[i].active {
						d.helpText = helpTexts[button.id]
					}
				}
			}
		}

		return game.SameState()
	}
}

func disableAllButtons(buttons []helpButton) {
	for i := range buttons {
		buttons[i].active = false
	}
}

type helpButton struct {
	id     helpButtonID
	x      int
	y      int
	active bool
}

func (button helpButton) isInside(x, y int) bool {
	return x >= button.x && x < button.x+16 && y >= button.y && y < button.y+16
}

type helpButtonID string

var helpTexts = map[helpButtonID]string{
	"customer":    "The customer wants pizza. They like some things (to the left), they hate some things (to the right).\nClick on the customer to give the pizza to them.",
	"reputation":  "Pizza that does not match the customer's expectations lowers your reputation. If you lose all your reputation, you'll get fired. Good pizza will raise your reputation.",
	"ingredients": "Ingrediens can be taken from here. You may put them back, unless they're already on the pizza.\nThe ingredient's flavors will be added to the pizza. You can see them when taking one.",
	"hotkeys":     "Hotkeys:\n\nR: Rotate\nShift+R: Rotate\nG: Toggle grid\nC: Pizza->Customer\nT: Leave help",
}
