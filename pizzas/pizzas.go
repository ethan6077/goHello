package pizzas

import "sort"

func GetPizzas() []string {
	var pizzas = []string{"Pepperoni", "Sausage", "Cheese", "Mushroom", "Pineapple"}
	pizzas = append(pizzas, "Ham")
	return pizzas
}

func SortPizzasTop3(pizzas []string) string {
	sort.Strings(pizzas)
	top3 := pizzas[0:3]
	top3String := ""
	for _, v := range top3 {
		top3String = top3String + v + "; "
	}
	return top3String
}

func GetPizzaMap() string {
	pizzaMap := make(map[string]string)
	pizzaMap["pep"] = "Pepperoni"
	pizzaMap["sau"] = "Sausage"
	pizzaMap["che"] = "Cheese"
	pizzaMap["mus"] = "Mushroom"
	pizzaMap["pin"] = "Pineapple"
	pizzaMap["ham"] = "Ham"

	pizzaString := ""
	for k, v := range pizzaMap {
		pizzaString = pizzaString + k + ": " + v + "; "
	}

	return pizzaString
}
