package game

import (
	"math/rand"

	"github.com/GodsBoss/gggg/v2/pkg/vector/vector2d"
)

type orders map[string]order

func (os orders) randomOrder(d *data) {
	keys := make([]string, 0)

	for key := range os {
		keys = append(keys, key)
	}

	key := keys[rand.Intn(len(keys))]
	os[key].applyTo(d)
}

var ingredientPositions = []vector2d.Vector[int]{
	vector2d.Cartesian[int](45, 170),
	vector2d.Cartesian[int](100, 180),
	vector2d.Cartesian[int](160, 185),
	vector2d.Cartesian[int](220, 180),
	vector2d.Cartesian[int](275, 170),
}

type order struct {
	// pizzaDiameter is the pizza diameter for this order.
	pizzaDiameter int

	leaveCorners bool

	// likes are the flavors the customer desires. If there are at least two preferred flavors, there's a chance one will be deleted.
	likes map[flavor]int

	// dislikes are the flavors the customer does not like. If there are at least two disliked flavors, there's a chance one will be deleted.
	dislikes map[flavor]struct{}

	// ingredients are the amounts of ingredients. Positions are automatically determined. There is a chance to get additional ingredients.
	ingredients map[ingredientType]int

	// fixedForgiveness is the base forgiveness value for the customer corresponding to this order.
	fixedForgiveness int

	// randomForgiveness is additional forgiveness that's added at random to the customer's forgiveness.
	randomForgiveness int
}

func (o order) applyTo(d *data) {
	// Create pizza.
	d.pizza = createPizza(o.pizzaDiameter, o.leaveCorners)

	// Create and configure a new customer.
	d.customer = &customer{
		likes:       make(map[flavor]int),
		dislikes:    make(map[flavor]struct{}),
		forgiveness: o.fixedForgiveness,
		mood:        customerMoodNormal,
		activity:    customerWaiting,
	}
	for fl := range o.likes {
		d.customer.likes[fl] = o.likes[fl]
	}
	for fl := range o.dislikes {
		d.customer.dislikes[fl] = o.dislikes[fl]
	}

	// Add a bit of randomness to the customer.

	if o.randomForgiveness > 0 {
		d.customer.forgiveness += rand.Intn(o.randomForgiveness + 1)
	}

	if len(d.customer.likes) >= 2 {
		for fl := range d.customer.likes {
			if rand.Intn(5) == 0 {
				delete(d.customer.likes, fl)
			}
			break
		}
	}

	if len(d.customer.dislikes) >= 2 {
		for fl := range d.customer.dislikes {
			if rand.Intn(5) == 0 {
				delete(d.customer.dislikes, fl)
			}
			break
		}
	}

	d.waitingIngredients = make([]waitingIngredient, 0)

	positions := make([]vector2d.Vector[int], len(ingredientPositions))
	copy(positions, ingredientPositions)
	rand.Shuffle(
		len(positions),
		func(i, j int) {
			positions[i], positions[j] = positions[j], positions[i]
		},
	)

	positionIndex := 0
	for typ, amount := range o.ingredients {
		if rand.Intn(5) == 0 {
			amount++
		}
		d.waitingIngredients = append(
			d.waitingIngredients,
			waitingIngredient{
				typ:    typ,
				amount: amount,
				x:      positions[positionIndex].X() - ingredientSizes[typ].Width/2,
				y:      positions[positionIndex].Y() - ingredientSizes[typ].Height,
			},
		)
		positionIndex++
	}
}

var possibleOrders = orders{
	"3_1": {
		pizzaDiameter: 3,
		likes: map[flavor]int{
			flavorSweet: 1,
			flavorFish:  2,
		},
		dislikes: map[flavor]struct{}{
			flavorCalamari: struct{}{},
		},
		ingredients: map[ingredientType]int{
			ingredientAnanas:      1,
			ingredientAnchovi:     2,
			ingredientRubberBoots: 1,
		},
		fixedForgiveness:  0,
		randomForgiveness: 0,
	},
	"4_1": {
		pizzaDiameter: 4,
		likes: map[flavor]int{
			flavorCalamari: 2,
			flavorSalty:    2,
		},
		dislikes: map[flavor]struct{}{
			flavorSweet: struct{}{},
		},
		ingredients: map[ingredientType]int{
			ingredientAnchovi:     2,
			ingredientRubberBoots: 4,
		},
		fixedForgiveness:  0,
		randomForgiveness: 0,
	},
	"4_2": {
		pizzaDiameter: 4,
		leaveCorners:  true,
		likes: map[flavor]int{
			flavorSweet:    1,
			flavorCalamari: 1,
			flavorFish:     4,
		},
		dislikes: map[flavor]struct{}{
			flavorFungus: struct{}{},
			flavorLiquid: struct{}{},
		},
		ingredients: map[ingredientType]int{
			ingredientAnanas:      1,
			ingredientAnchovi:     3,
			ingredientSquid:       2,
			ingredientTomatoSauce: 1,
			ingredientMushroom:    1,
		},
	},
	"5_1": {
		pizzaDiameter: 5,
		likes: map[flavor]int{
			flavorCalamari: 2,
		},
		dislikes: map[flavor]struct{}{
			flavorSweet: struct{}{},
			flavorFish:  struct{}{},
		},
		ingredients: map[ingredientType]int{
			ingredientAnanas:      1,
			ingredientAnchovi:     1,
			ingredientRubberBoots: 2,
		},
		fixedForgiveness:  2,
		randomForgiveness: 1,
	},
	"5_2": {
		pizzaDiameter: 5,
		likes: map[flavor]int{
			flavorMeat: 4,
		},
		ingredients: map[ingredientType]int{
			ingredientBurger: 4,
			ingredientSalami: 2,
		},
	},
}
