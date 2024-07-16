package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")

	// message := greetings.Greet("Ethan")

	// fmt.Println(message)

	// n := time.Now()
	// fmt.Println(n)

	// cmd.StartCli()

	// myPizzas := pizzas.GetPizzas()
	// top3 := pizzas.SortPizzasTop3(myPizzas)
	// fmt.Println(top3)

	// pizzaMap := pizzas.GetPizzaMap()
	// fmt.Println(pizzaMap)

	p := Person{Name: "Ethan", countGreets: 0}
	fmt.Println(p.Greet())
	fmt.Println(p.Greet())
	fmt.Println(p.Greet())
	fmt.Println(p.Bye())

}
