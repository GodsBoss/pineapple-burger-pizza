package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
	"github.com/GodsBoss/gggg/v2/pkg/game"
	"github.com/GodsBoss/gggg/v2/pkg/vector/vector2d"
)

const playingState = "playing"

func initPlaying(d *data) game.NextState {
	d.state = playingState
	d.score = 0
	d.pizzaGridOverlayVisible = false
	d.reputation = 10
	getNewOrder(d)

	return game.SameState()
}

func createReceiveKeyEventPlaying() func(d *data, event keyboard.Event) game.NextState {
	return func(d *data, event keyboard.Event) game.NextState {
		if event.Key == "g" && keyboard.IsUpEvent(event) {
			d.pizzaGridOverlayVisible = !d.pizzaGridOverlayVisible
		}

		if event.Key == "r" && keyboard.IsDownEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.orientation = (d.draggedIngredient.orientation + ingredientClockwise) % 4
			d.draggedIngredient.fields = rotateFields(d.draggedIngredient.fields, clockwise)
			calculateIngredientTargetFields(d)
		}

		if event.Key == "R" && keyboard.IsDownEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.orientation = (d.draggedIngredient.orientation + ingredientCounterClockwise) % 4
			d.draggedIngredient.fields = rotateFields(d.draggedIngredient.fields, counterClockwise)
			calculateIngredientTargetFields(d)
		}

		// Give pizza to customer.
		if event.Key == "c" && keyboard.IsDownEvent(event) {
			customerGetsPizza(d)
			getNewOrder(d)
		}

		return game.SameState()
	}
}

func createReceiveMouseEventPlaying() func(d *data, event mouse.Event) game.NextState {
	return func(d *data, event mouse.Event) game.NextState {
		if mouse.IsPrimaryButtonEvent(event) && mouse.IsDownEvent(event) {
			if d.draggedIngredient != nil && d.draggedIngredient.isValidPlacement() {
				placeIngredient(d)
				return game.SameState()
			}

			for i, ingredient := range d.waitingIngredients {
				if ingredient.inside(event.X, event.Y) {

					if d.draggedIngredient == nil && ingredient.amount > 0 {
						d.draggedIngredient = &draggedIngredient{
							typ:         ingredient.typ,
							orientation: ingredientUp,
							x:           event.X,
							y:           event.Y,
							fields:      make([]vector2d.Vector[int], len(ingredientFields[ingredient.typ])),
						}
						copy(d.draggedIngredient.fields, ingredientFields[ingredient.typ])
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

		if mouse.IsPrimaryButtonEvent(event) && mouse.IsUpEvent(event) {
			if event.X > 10 && event.X < 100 && event.Y > 0 && event.Y < 50 {
				customerGetsPizza(d)
				getNewOrder(d)
			}
		}

		if mouse.IsMoveEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.x = event.X
			d.draggedIngredient.y = event.Y

			calculateIngredientTargetFields(d)
		}

		return game.SameState()
	}
}

func createReceiveTickEventPlaying(gameOverState game.StateID) func(d *data, event tick.Event) game.NextState {
	return func(d *data, event tick.Event) game.NextState {
		if d.reputation <= 0 {
			return game.SwitchState(gameOverState)
		}
		return game.SameState()
	}
}

func placeIngredient(d *data) {
	for _, field := range d.draggedIngredient.validFields {
		d.pizza.grid[field.X()][field.Y()].occupied = true
	}
	for flavor, amount := range ingredientFlavors[d.draggedIngredient.typ] {
		d.pizza.flavors[flavor] += amount
	}
	d.placedIngredients = append(
		d.placedIngredients,
		placedIngredient{
			typ:         d.draggedIngredient.typ,
			orientation: d.draggedIngredient.orientation,
			x:           d.draggedIngredient.x,
			y:           d.draggedIngredient.y,
			width:       d.draggedIngredient.Width(),
			height:      d.draggedIngredient.Height(),
		},
	)
	d.score += scoreForPlacedIngredient(*d.draggedIngredient)
	d.draggedIngredient = nil
}

func getNewOrder(d *data) {
	d.placedIngredients = make([]placedIngredient, 0)
	d.draggedIngredient = nil
	possibleOrders.randomOrder(d)
}

func calculateIngredientTargetFields(d *data) {
	d.draggedIngredient.validFields = make([]vector2d.Vector[int], 0)
	d.draggedIngredient.invalidFields = make([]vector2d.Vector[int], 0)

	atLeastOneFieldInPizzaBounds := false

	for _, field := range d.draggedIngredient.fields {
		fieldOffsetX := field.X()*pizzaFieldWidth + d.draggedIngredient.x - 160 + d.pizza.Width()*pizzaFieldWidth/2 - d.draggedIngredient.Width()/2 + pizzaFieldWidth/2
		fieldOffsetY := field.Y()*pizzaFieldHeight + d.draggedIngredient.y - 100 + d.pizza.Height()*pizzaFieldHeight/2 - d.draggedIngredient.Height()/2 + pizzaFieldHeight/2

		// We need to shift the result for negative results, because rounding has a bias towards 0. E.g. every value from -15 to 15 will be
		// rounded towards 0 if divided by 16.

		if fieldOffsetX < 0 {
			fieldOffsetX -= pizzaFieldWidth - 1
		}
		if fieldOffsetY < 0 {
			fieldOffsetY -= pizzaFieldHeight - 1
		}

		fieldOffsetX, fieldOffsetY = fieldOffsetX/pizzaFieldWidth, fieldOffsetY/pizzaFieldHeight

		withinPizzaBounds := fieldOffsetX >= 0 && fieldOffsetX < d.pizza.Width() && fieldOffsetY >= 0 && fieldOffsetY < d.pizza.Height()
		if withinPizzaBounds && !d.pizza.grid[fieldOffsetX][fieldOffsetY].invalid {
			atLeastOneFieldInPizzaBounds = true
			if d.pizza.grid[fieldOffsetX][fieldOffsetY].occupied {
				d.draggedIngredient.invalidFields = append(d.draggedIngredient.invalidFields, vector2d.Cartesian[int](fieldOffsetX, fieldOffsetY))
			} else {
				d.draggedIngredient.validFields = append(d.draggedIngredient.validFields, vector2d.Cartesian[int](fieldOffsetX, fieldOffsetY))
			}
		} else {
			d.draggedIngredient.invalidFields = append(d.draggedIngredient.invalidFields, vector2d.Cartesian[int](fieldOffsetX, fieldOffsetY))
		}
	}

	if !atLeastOneFieldInPizzaBounds {
		d.draggedIngredient.validFields = make([]vector2d.Vector[int], 0)
		d.draggedIngredient.invalidFields = make([]vector2d.Vector[int], 0)
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

	// Disable corner fields only if diameter > 3.
	if diameter > 3 {
		grid[0][0].invalid = true
		grid[max][0].invalid = true
		grid[0][max].invalid = true
		grid[max][max].invalid = true
	}

	return &pizza{
		grid:    grid,
		flavors: make(map[flavor]int),
	}
}

type pizza struct {
	grid    [][]pizzaField
	flavors map[flavor]int
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
