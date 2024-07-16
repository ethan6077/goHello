package main

import "strconv"

type Person struct {
	Name        string
	countGreets int
}

func (p *Person) Greet() string {
	p.countGreets++
	return "Hello, I am " + p.Name
}

func (p Person) Bye() string {
	greetTimes := strconv.Itoa(p.countGreets)
	return "Bye, I am " + p.Name + ". I have greeted " + greetTimes + " times."
}
