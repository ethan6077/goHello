package main

import (
	"fmt"
	"time"

	"ethan.com/greetings"
	"ethan.com/pizzas"
)

func main() {
	fmt.Println("Hello, world.")

	message := greetings.Greet("Ethan")

	fmt.Println(message)

	n := time.Now()
	fmt.Println(n)

	// cmd.StartCli()

	myPizzas := pizzas.GetPizzas()
	top3 := pizzas.SortPizzasTop3(myPizzas)
	fmt.Println(top3)

	pizzaMap := pizzas.GetPizzaMap()
	fmt.Println(pizzaMap)

}
