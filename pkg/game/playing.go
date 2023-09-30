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
		{
			typ:    ingredientAnanas,
			amount: 2,
			x:      200,
			y:      130,
		},
		{
			typ:    ingredientRubberBoots,
			amount: 3,
			x:      80,
			y:      160,
		},
	}

	d.customer = &customer{
		likes: map[flavor]int{
			flavorCalamari: 2,
		},
		dislikes: map[flavor]struct{}{
			flavorSweet: struct{}{},
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
			d.draggedIngredient.fields = rotateFields(d.draggedIngredient.fields, clockwise)
			calculateIngredientTargetFields(d)
		}

		if event.Key == "R" && keyboard.IsDownEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.orientation = (d.draggedIngredient.orientation + ingredientCounterClockwise) % 4
			d.draggedIngredient.fields = rotateFields(d.draggedIngredient.fields, counterClockwise)
			calculateIngredientTargetFields(d)
		}

		return game.SameState()
	}
}

func createReceiveMouseEventPlaying() func(d *data, event mouse.Event) game.NextState {
	return func(d *data, event mouse.Event) game.NextState {
		if mouse.IsPrimaryButtonEvent(event) && mouse.IsDownEvent(event) {
			if d.draggedIngredient != nil && d.draggedIngredient.isValidPlacement() {
				for _, field := range d.draggedIngredient.validFields {
					d.pizza.grid[field.X()][field.Y()].occupied = true
				}
				for flavor, amount := range ingredientFlavors[d.draggedIngredient.typ] {
					d.pizza.flavors[flavor] += amount
				}
				d.draggedIngredient = nil
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

		if mouse.IsMoveEvent(event) && d.draggedIngredient != nil {
			d.draggedIngredient.x = event.X
			d.draggedIngredient.y = event.Y

			calculateIngredientTargetFields(d)
		}

		return game.SameState()
	}
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

	// Disable corner fields.
	grid[0][0].invalid = true
	grid[max][0].invalid = true
	grid[0][max].invalid = true
	grid[max][max].invalid = true

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
	ingredientAnchovi     ingredientType = "anchovy"
	ingredientAnanas      ingredientType = "ananas"
	ingredientRubberBoots ingredientType = "rubber_boots"
)

// ingredientSizes are the sizes for waiting ingredients.
var ingredientSizes = func(fieldsPerIngredientType map[ingredientType][]vector2d.Vector[int]) map[ingredientType]size {

	sizes := make(map[ingredientType]size)

	for key, fields := range fieldsPerIngredientType {
		maxX, maxY := 0, 0

		for _, field := range fields {
			if x := field.X(); x > maxX {
				maxX = x
			}
			if y := field.Y(); y > maxY {
				maxY = y
			}
		}

		sizes[key] = size{
			Width:  (maxX + 1) * pizzaFieldWidth,
			Height: (maxY + 1) * pizzaFieldHeight,
		}
	}

	return sizes

}(ingredientFields)

// ingredientFields maps ingredient types to the fields it will occupy on a pizza.
// This refers to the fields when in "up" orientation, so when ingredients are
// rotated, the fields are rotated as well.
var ingredientFields = map[ingredientType][]vector2d.Vector[int]{
	ingredientAnchovi: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](1, 0),
	},
	ingredientAnanas: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](0, 2),
		vector2d.Cartesian[int](1, 1),
		vector2d.Cartesian[int](1, 2),
		vector2d.Cartesian[int](2, 0),
	},
	ingredientRubberBoots: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](1, 1),
	},
}

// rotateFields rotates the given fields. Fields are left/top aligned. The left-most
// fields always have an X coordinate of 0, the top-most fields have an Y coordinate of 0.
func rotateFields(fields []vector2d.Vector[int], rotate func(int, int) (int, int)) []vector2d.Vector[int] {
	// smallestX and smallestY will be used to determine the offset the intermediate list needs to be shifted.
	// As a list of fields has at least one field with X = 0 and at least one field with Y = 0, it is safe to
	// assume that the smallest X and Y cannot be greater than 0.
	// These numbers will always be zero or less.
	smallestX, smallestY := 0, 0

	result := make([]vector2d.Vector[int], len(fields))

	for i := range fields {
		// New coordinates.
		x, y := rotate(fields[i].X(), fields[i].Y())

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

// clockwise rotates x and y clockwise.
func clockwise(x int, y int) (int, int) {
	return -y, x
}

// counterClockwise rotates x and y counter-clockwise.
func counterClockwise(x int, y int) (int, int) {
	return y, -x
}

type draggedIngredient struct {
	typ           ingredientType
	orientation   ingredientOrientation
	x             int
	y             int
	fields        []vector2d.Vector[int]
	validFields   []vector2d.Vector[int]
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

func (ingr draggedIngredient) isValidPlacement() bool {
	return len(ingr.fields) == len(ingr.validFields) && len(ingr.invalidFields) == 0
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

type flavor string

const (
	flavorSweet    flavor = "sweet"
	flavorCalamari flavor = "calamari"
	flavorSalty    flavor = "salty"
	flavorFish     flavor = "fish"
)

// flavorList provides a consistent sort order for flavors. This is useful as maps don't have an inherent sort order.
var flavorList = []flavor{
	flavorCalamari,
	flavorFish,
	flavorSalty,
	flavorSweet,
}

var ingredientFlavors = map[ingredientType]map[flavor]int{
	ingredientAnanas: {
		flavorSweet: 1,
	},
	ingredientAnchovi: {
		flavorFish:  1,
		flavorSalty: 1,
	},
	ingredientRubberBoots: {
		flavorCalamari: 1,
	},
}

type customer struct {
	likes    map[flavor]int
	dislikes map[flavor]struct{}
}

// ratePizza lets the customer rate a pizza. The best possible rating is 0. Usually, ratings are negative.
func (c customer) ratePizza(p pizza) int {
	rating := 0

	for fl, l := range c.likes {
		if p.flavors[fl] < l {
			rating -= l - p.flavors[fl]
		}
	}

	for fl, _ := range c.dislikes {
		rating -= p.flavors[fl]
	}

	return rating
}
