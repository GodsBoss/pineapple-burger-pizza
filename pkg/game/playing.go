package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

const playingState = "playing"

func initPlaying(d *data) game.NextState {
	d.state = playingState
	d.pizzaGridOverlayVisible = false
	d.pizza = createPizza(5)
	d.waitingIngredients = []waitingIngredient{
		{
			typ:    ingredientAnchovi,
			amount: 5,
			x:      20,
			y:      160,
		},
	}

	return game.SameState()
}

func createReceiveKeyEventPlaying() func(d *data, event keyboard.Event) game.NextState {
	return func(d *data, event keyboard.Event) game.NextState {
		if event.Key == "g" && keyboard.IsUpEvent(event) {
			d.pizzaGridOverlayVisible = !d.pizzaGridOverlayVisible
		}

		if event.Key == "r" && keyboard.IsDownEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.orientation = (d.draggedIngredient.orientation + ingredientClockwise) % 4
		}

		if event.Key == "R" && keyboard.IsDownEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.orientation = (d.draggedIngredient.orientation + ingredientCounterClockwise) % 4
		}

		return game.SameState()
	}
}

func createReceiveMouseEventPlaying() func(d *data, event mouse.Event) game.NextState {
	return func(d *data, event mouse.Event) game.NextState {
		if mouse.IsPrimaryButtonEvent(event) && mouse.IsDownEvent(event) {
			for i, ingredient := range d.waitingIngredients {
				if ingredient.inside(event.X, event.Y) {

					if d.draggedIngredient == nil && ingredient.amount > 0 {
						d.draggedIngredient = &draggedIngredient{
							typ:         ingredient.typ,
							orientation: ingredientUp,
							x:           event.X,
							y:           event.Y,
						}
						d.waitingIngredients[i].amount--
						return game.SameState()
					}

					if d.draggedIngredient != nil && ingredient.typ == d.draggedIngredient.typ {
						d.waitingIngredients[i].amount++
						d.draggedIngredient = nil
						return game.SameState()
					}
				}
			}
		}

		if mouse.IsMoveEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.x = event.X
			d.draggedIngredient.y = event.Y
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

// waitingIngredient is an ingredient waiting on the table.
type waitingIngredient struct {
	typ    ingredientType
	amount int
	x      int
	y      int
}

func (ingr waitingIngredient) inside(x int, y int) bool {
	s, ok := ingredientSizes[ingr.typ]
	if !ok {
		return false
	}
	return x >= ingr.x && y >= ingr.y && x <= ingr.x+s.Width && y <= ingr.y+s.Height
}

type ingredientType string

const (
	ingredientAnchovi ingredientType = "anchovy"
)

// ingredientSizes are the sizes for waiting ingredients.
var ingredientSizes = map[ingredientType]size{
	ingredientAnchovi: size{
		Width:  40,
		Height: 16,
	},
}

type draggedIngredient struct {
	typ         ingredientType
	orientation ingredientOrientation
	x           int
	y           int
}

// ingredientOrientation determines whether an ingredient is up, down, etc.
type ingredientOrientation int

const (
	ingredientUp    ingredientOrientation = 0
	ingredientRight ingredientOrientation = 1
	ingredientDown  ingredientOrientation = 2
	ingredientLeft  ingredientOrientation = 3
)

const (
	ingredientClockwise        ingredientOrientation = 1
	ingredientCounterClockwise ingredientOrientation = 3
)
