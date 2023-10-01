package game

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
