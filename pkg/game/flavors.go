package game

type flavor string

const (
	flavorCalamari flavor = "calamari"
	flavorFish     flavor = "fish"
	flavorMeat     flavor = "meat"
	flavorSalty    flavor = "salty"
	flavorSweet    flavor = "sweet"
)

// flavorList provides a consistent sort order for flavors. This is useful as maps don't have an inherent sort order.
var flavorList = []flavor{
	flavorCalamari,
	flavorFish,
	flavorMeat,
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
