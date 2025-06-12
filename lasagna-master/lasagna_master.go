package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prep_time int) int {
	if prep_time == 0 {
		prep_time = 2
	}
	return len(layers) * prep_time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodle_count := 0
	sauce_count := 0.0
	for _, layer := range layers {
		if layer == "sauce" {
			sauce_count++
			continue
		}
		if layer == "noodles" {
			noodle_count++
		}
	}
	return 50 * noodle_count, 0.2 * sauce_count
}

// Function `AddSecretIngredient` that accepts two slices of ingredients of type `[]string` as parameters.
// The first parameter is the list your friend sent you, the second is the ingredient list of your own recipe.
// The last element in your ingredient list is always `"?"`.
// The function should replace it with the last item from your friends list.
func AddSecretIngredient(friendsList, myList []string) {
	if myList[len(myList)-1] == "?" {
		myList[len(myList)-1] = friendsList[len(friendsList)-1]
	}
}

// ScaleRecipe` that takes two parameters.
// - A slice of `float64` amounts needed for 2 portions.
// - The number of portions you want to cook.
// The function should return a slice of `float64` of the amounts needed for the desired number of portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	ratio := float64(portions) / 2.0
	out := make([]float64, len(quantities))
	for i, qty := range quantities {
		out[i] = ratio * qty
	}
	return out
}
