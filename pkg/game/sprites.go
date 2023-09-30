package game

import (
	"strconv"

	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
	r "github.com/GodsBoss/gggg/v2/pkg/rendering/canvas2drendering"
)

func createSpriteMap(sourceImage *dom.Image) (canvas2drendering.SpriteMap, spriteKeys) {
	spriteMap := canvas2drendering.NewSpriteMap(sourceImage)
	keys := spriteKeys{}
	addSprite := createAddSprite(spriteMap)

	keys.backgroundTitle = addSprite("background_title")
	keys.backgroundPlaying = addSprite("background_playing")
	keys.backgroundGameOver = addSprite("background_game_over")

	keys.pizzaGridOverlayFree = addSprite("pizza_grid_overlay_free")
	keys.pizzaGridOverlayOccupied = addSprite("pizza_grid_overlay_occupied")

	keys.ingredientGridOverlayFree = addSprite("ingredient_grid_overlay_free")
	keys.ingredientGridOverlayOccupied = addSprite("ingredient_grid_overlay_occupied")

	addIngredientSprites := createAddIngredientSprites(spriteMap)

	keys.ingredientAnchovi = addIngredientSprites("anchovy")
	keys.ingredientAnanas = addIngredientSprites("ananas")
	keys.ingredientRubberBoots = addIngredientSprites("rubber_boots")

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
	backgroundTitle    r.SpriteKey
	backgroundPlaying  r.SpriteKey
	backgroundGameOver r.SpriteKey

	pizzaGridOverlayFree     r.SpriteKey
	pizzaGridOverlayOccupied r.SpriteKey

	ingredientGridOverlayFree     r.SpriteKey
	ingredientGridOverlayOccupied r.SpriteKey

	ingredientAnchovi     [4]r.SpriteKey
	ingredientAnanas      [4]r.SpriteKey
	ingredientRubberBoots [4]r.SpriteKey
}
