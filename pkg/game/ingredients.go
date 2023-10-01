package game

import "github.com/GodsBoss/gggg/v2/pkg/vector/vector2d"

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
	ingredientAnanas      ingredientType = "ananas"
	ingredientAnchovi     ingredientType = "anchovy"
	ingredientBurger      ingredientType = "burger"
	ingredientMushroom    ingredientType = "mushroom"
	ingredientRubberBoots ingredientType = "rubber_boots"
	ingredientSalami      ingredientType = "salami"
	ingredientSquid       ingredientType = "squid"
	ingredientTomatoSauce ingredientType = "tomato_sauce"
)

var ingredientList = []ingredientType{
	ingredientAnanas,
	ingredientAnchovi,
	ingredientBurger,
	ingredientMushroom,
	ingredientRubberBoots,
	ingredientSalami,
	ingredientSquid,
	ingredientTomatoSauce,
}

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
	ingredientAnanas: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](0, 2),
		vector2d.Cartesian[int](1, 1),
		vector2d.Cartesian[int](1, 2),
		vector2d.Cartesian[int](2, 0),
	},
	ingredientAnchovi: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](1, 0),
	},
	ingredientBurger: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](1, 0),
		vector2d.Cartesian[int](2, 0),
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](1, 1),
		vector2d.Cartesian[int](2, 1),
	},
	ingredientMushroom: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](1, 0),
		vector2d.Cartesian[int](1, 1),
	},
	ingredientRubberBoots: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](1, 1),
	},
	ingredientSalami: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](1, 0),
		vector2d.Cartesian[int](2, 1),
	},
	ingredientSquid: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](0, 2),
		vector2d.Cartesian[int](1, 0),
		vector2d.Cartesian[int](1, 1),
	},
	ingredientTomatoSauce: []vector2d.Vector[int]{
		vector2d.Cartesian[int](0, 0),
		vector2d.Cartesian[int](0, 1),
		vector2d.Cartesian[int](1, 0),
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

type placedIngredient struct {
	typ         ingredientType
	orientation ingredientOrientation
	x           int
	y           int
	width       int
	height      int
}
