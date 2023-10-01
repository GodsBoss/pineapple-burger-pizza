package game

// data represents all of the game's data.
type data struct {
	// state is the state the game is in (title, playing, etc.). Set by states, used for rendering.
	state string

	// pizza is the pizza visible when playing.
	pizza *pizza

	// pizzaGridOverlayVisible determines whether the pizza grid overlay is visible.
	pizzaGridOverlayVisible bool

	// waitingIngredients are the ingredients that the player can take.
	waitingIngredients []waitingIngredient

	// draggedIngredient is the ingredient currently dragged by the player.
	draggedIngredient *draggedIngredient

	// placedIngredients are the ingredients already placed on the pizza.
	placedIngredients []placedIngredient

	// customer is the customer waiting for the pizza.
	customer *customer

	// reputation is your reputation as a pizza baker. If this reaches zero, the game is over.
	reputation int

	// score is the score of this run.
	score int

	// highscore is the highest score achieved.
	highscore int

	// helpButtons are the help buttons found on the help page.
	helpButtons []helpButton

	// helpText is the text shown as help.
	helpText string
}
