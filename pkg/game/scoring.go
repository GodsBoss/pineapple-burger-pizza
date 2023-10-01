package game

func scoreForPlacedIngredient(ingr draggedIngredient) int {
	return len(ingr.fields) * 5
}

func scoreForRating(rating int) int {
	if rating < 0 {
		return 10
	}

	return 10 + rating*10
}
