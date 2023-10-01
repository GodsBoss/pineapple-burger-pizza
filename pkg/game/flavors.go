package game

type flavor string

const (
	flavorCalamari flavor = "calamari"
	flavorFish     flavor = "fish"
	flavorLiquid   flavor = "liquid"
	flavorMeat     flavor = "meat"
	flavorFungus   flavor = "fungus"
	flavorSalty    flavor = "salty"
	flavorSweet    flavor = "sweet"
)

// flavorList provides a consistent sort order for flavors. This is useful as maps don't have an inherent sort order.
var flavorList = []flavor{
	flavorCalamari,
	flavorFish,
	flavorLiquid,
	flavorMeat,
	flavorFungus,
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
	ingredientBurger: {
		flavorMeat: 1,
	},
	ingredientMushroom: {
		flavorFungus: 1,
	},
	ingredientRubberBoots: {
		flavorCalamari: 1,
	},
	ingredientSalami: {
		flavorMeat:  1,
		flavorSalty: 1,
	},
	ingredientSquid: {
		flavorCalamari: 1,
		flavorFish:     1,
	},
	ingredientTomatoSauce: {
		flavorLiquid: 1,
	},
}
