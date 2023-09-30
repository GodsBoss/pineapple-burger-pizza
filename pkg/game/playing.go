package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/game"
	"github.com/GodsBoss/gggg/v2/pkg/vector/vector2d"
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
			d.draggedIngredient.fields = rotateFields(d.draggedIngredient.fields)
			calculateIngredientTargetFields(d)
		}

		if event.Key == "R" && keyboard.IsDownEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.orientation = (d.draggedIngredient.orientation + ingredientCounterClockwise) % 4
			d.draggedIngredient.fields = rotateFields(rotateFields(rotateFields(d.draggedIngredient.fields))) // Dirty hack! Use counter-clockwise rotation instead.
			calculateIngredientTargetFields(d)
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

		if mouse.IsMoveEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.x = event.X
			d.draggedIngredient.y = event.Y

			calculateIngredientTargetFields(d)
		}

		return game.SameState()
	}
}

func calculateIngredientTargetFields(d *data) {
	for col := range d.pizza.grid {
		for row := range d.pizza.grid[col] {
			d.pizza.grid[col][row].draggedIngredientTarget = false
		}
	}

	d.draggedIngredient.invalidFields = make([]vector2d.Vector[int], 0)

	atLeastOneFieldInPizzaBounds := false

	for _, field := range d.draggedIngredient.fields {
		fieldOffsetX := (field.X()*pizzaFieldWidth + d.draggedIngredient.x - 160 + d.pizza.Width()*pizzaFieldWidth/2 - d.draggedIngredient.Width()/2 + pizzaFieldWidth/2) / pizzaFieldWidth
		fieldOffsetY := (field.Y()*pizzaFieldHeight + d.draggedIngredient.y - 100 + d.pizza.Height()*pizzaFieldHeight/2 - d.draggedIngredient.Height()/2 + pizzaFieldHeight/2) / pizzaFieldHeight

		withinPizzaBounds := fieldOffsetX >= 0 && fieldOffsetX < d.pizza.Width() && fieldOffsetY >= 0 && fieldOffsetY < d.pizza.Height()
		if withinPizzaBounds && !d.pizza.grid[fieldOffsetX][fieldOffsetY].invalid {
			atLeastOneFieldInPizzaBounds = true
			d.pizza.grid[fieldOffsetX][fieldOffsetY].draggedIngredientTarget = true
		} else {
			d.draggedIngredient.invalidFields = append(d.draggedIngredient.invalidFields, vector2d.Cartesian[int](fieldOffsetX, fieldOffsetY))
		}

	}

	if !atLeastOneFieldInPizzaBounds {
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

	// draggedIngredientTarget is true if an ingredient's field is dragged over this field.
	draggedIngredientTarget bool
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
		Width:  pizzaFieldWidth * 2,
		Height: pizzaFieldHeight,
	},
}

// ingredientFields maps ingredient types to the fields it will occupy on a pizza.
// This refers to the fields when in "up" orientation, so when ingredients are
// rotated, the fields are rotated as well.
var ingredientFields = map[ingredientType][]vector2d.Vector[int]{
	ingredientAnchovi: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](1, 0),
	},
}

// rotateFields rotates the given fields. Fields are left/top aligned. The left-most
// fields always have an X coordinate of 0, the top-most fields have an Y coordinate of 0.
func rotateFields(fields []vector2d.Vector[int]) []vector2d.Vector[int] {
	// smallestX and smallestY will be used to determine the offset the intermediate list needs to be shifted.
	// As a list of fields has at least one field with X = 0 and at least one field with Y = 0, it is safe to
	// assume that the smallest X and Y cannot be greater than 0.
	// These numbers will always be zero or less.
	smallestX, smallestY := 0, 0

	result := make([]vector2d.Vector[int], len(fields))

	for i := range fields {
		// New coordinates.
		x, y := fields[i].Y(), -fields[i].X()

		if x < smallestX {
			smallestX = x
		}

		if y < smallestY {
			smallestY = y
		}

		result[i] = vector2d.Cartesian[int](x, y)
	}

	offset := vector2d.Cartesian[int](-smallestX, -smallestY)

	for i := range result {
		result[i] = vector2d.Sum[int](result[i], offset)
	}

	return result
}

type draggedIngredient struct {
	typ           ingredientType
	orientation   ingredientOrientation
	x             int
	y             int
	fields        []vector2d.Vector[int]
	invalidFields []vector2d.Vector[int]
}

func (ingr draggedIngredient) Width() int {
	maxX := 0
	for _, field := range ingr.fields {
		if x := field.X(); x > maxX {
			maxX = x
		}
	}
	return (maxX + 1) * pizzaFieldWidth
}

func (ingr draggedIngredient) Height() int {
	maxY := 0
	for _, field := range ingr.fields {
		if y := field.Y(); y > maxY {
			maxY = y
		}
	}
	return (maxY + 1) * pizzaFieldHeight
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
