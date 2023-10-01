package game

import (
	"strconv"

	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
	r "github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
)

func createSpriteMap(sourceImage *dom.Image) (canvas2drendering.SpriteMap, spriteKeys) {
	spriteMap := canvas2drendering.NewSpriteMap(sourceImage)
	keys := spriteKeys{
		ingredients: make(map[ingredientType][4]r.SpriteKey),
		flavors:     make(map[flavor]r.SpriteKey),
		pizzas:      make(map[int]r.SpriteKey),
	}

	addSprite := createAddSprite(spriteMap)

	keys.backgroundTitle = addSprite("background_title")
	keys.backgroundPlaying = addSprite("background_playing")
	keys.backgroundTable = addSprite("background_table")

	keys.pizzaGridOverlayFree = addSprite("pizza_grid_overlay_free")
	keys.pizzaGridOverlayOccupied = addSprite("pizza_grid_overlay_occupied")

	keys.ingredientGridOverlayFree = addSprite("ingredient_grid_overlay_free")
	keys.ingredientGridOverlayOccupied = addSprite("ingredient_grid_overlay_occupied")

	keys.customerLike = addSprite("customer_like")
	keys.customerDislike = addSprite("customer_dislike")

	keys.reputationOK = addSprite("reputation_ok")
	keys.reputationGone = addSprite("reputation_gone")

	keys.customerHeadNormal = addSprite("customer_head_normal")
	keys.customerHeadAngry = addSprite("customer_head_angry")
	keys.customerHeadHappy = addSprite("customer_head_happy")
	keys.customerHeadEating = addSprite("customer_head_eating")
	keys.customerBody = addSprite("customer_body")

	keys.helpIcon = addSprite("help_icon")
	keys.helpIconActive = addSprite("help_icon_active")

	keys.keyboard = addSprite("keyboard")

	addIngredientSprites := createAddIngredientSprites(spriteMap)

	ingredientKeys := []ingredientType{
		ingredientAnanas,
		ingredientAnchovi,
		ingredientRubberBoots,
	}

	for _, ingredientKey := range ingredientKeys {
		keys.ingredients[ingredientKey] = addIngredientSprites(string(ingredientKey))
	}

	flavors := []flavor{
		flavorCalamari,
		flavorFish,
		flavorSalty,
		flavorSweet,
	}

	for _, fl := range flavors {
		keys.flavors[fl] = addSprite("flavor_" + string(fl))
	}

	for p := 3; p <= 5; p++ {
		keys.pizzas[p] = addSprite("pizza_" + strconv.Itoa(p))
	}

	return spriteMap, keys
}

func createAddSprite(spriteMap canvas2drendering.SpriteMap) func(key string) canvas2drendering.SpriteKey {
	return func(key string) canvas2drendering.SpriteKey {
		return spriteMap.AddSpriteSpecification(
			map[canvas2drendering.SpriteAttributes]canvas2drendering.SpriteData{
				canvas2drendering.SpriteAttributes{}: spritesData[key],
			},
		)
	}
}

func createAddIngredientSprites(spriteMap canvas2drendering.SpriteMap) func(key string) [4]canvas2drendering.SpriteKey {
	return func(key string) [4]canvas2drendering.SpriteKey {
		spriteKeys := [4]canvas2drendering.SpriteKey{}

		for suffix := 0; suffix <= 3; suffix++ {
			spriteKeys[suffix] = spriteMap.AddSpriteSpecification(
				map[canvas2drendering.SpriteAttributes]canvas2drendering.SpriteData{
					canvas2drendering.SpriteAttributes{}: spritesData[key+"_"+strconv.Itoa(suffix)],
				},
			)
		}

		return spriteKeys
	}
}

// spriteKeys holds the sprite keys generated when adding sprite specs to the spriteKeys factory.
type spriteKeys struct {
	backgroundTitle   r.SpriteKey
	backgroundPlaying r.SpriteKey
	backgroundTable   r.SpriteKey

	pizzaGridOverlayFree     r.SpriteKey
	pizzaGridOverlayOccupied r.SpriteKey

	ingredientGridOverlayFree     r.SpriteKey
	ingredientGridOverlayOccupied r.SpriteKey

	customerLike    r.SpriteKey
	customerDislike r.SpriteKey

	reputationOK   r.SpriteKey
	reputationGone r.SpriteKey

	customerHeadNormal r.SpriteKey
	customerHeadAngry  r.SpriteKey
	customerHeadHappy  r.SpriteKey
	customerHeadEating r.SpriteKey
	customerBody       r.SpriteKey

	helpIcon       r.SpriteKey
	helpIconActive r.SpriteKey

	keyboard r.SpriteKey

	ingredients map[ingredientType][4]r.SpriteKey
	flavors     map[flavor]r.SpriteKey
	pizzas      map[int]r.SpriteKey
}
