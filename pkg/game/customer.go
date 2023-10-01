package game

type customer struct {
	// likes determines what the customer expects from their pizza. Every missing flavor lets the rating suffer.
	likes map[flavor]int

	// dislikes determines what the customer does not want in their pizza. If disliked flavors are added, the rating suffers.
	dislikes map[flavor]struct{}

	// forgiveness is the customer's tolerance for bad pizza.
	forgiveness int

	// mood is the customer's mood. This is influenced
	mood customerMood

	// activity is the customer's current activity.
	activity customerActivity

	// remainingActivityTime is the remaining time for activities in ms.
	remainingActivityTime int
}

type customerMood string

const (
	customerMoodNormal customerMood = "normal"
	customerMoodAngry  customerMood = "angry"
	customerMoodHappy  customerMood = "happy"
)

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

func customerGetsPizza(d *data) {
	d.customer.activity = customerEating
	d.customer.remainingActivityTime = 2000

	rating := d.customer.ratePizza(*d.pizza) + d.customer.forgiveness
	if rating > 0 {
		d.reputation++
		if d.reputation > 10 {
			d.reputation = 10
		}
	}
	if rating < 0 { // Bad pizza.
		d.reputation--
	}
	if rating < -5 { // Very bad pizza.
		d.reputation--
	}
	if d.reputation < 0 {
		d.reputation = 0
	}
	d.score += scoreForRating(rating)
}

type customerActivity string

const (
	// customerWaiting is the customer's activity when waiting for the pizza.
	customerWaiting customerActivity = "waiting"

	// customerEating is the customer eating.
	customerEating customerActivity = "eating"

	// customerExperiencing is the customer experiencing the flavors of the pizza.
	customerExperiencing customerActivity = "experiencing"
)
