module ethan.com/hello

go 1.22.5

replace ethan.com/greetings => ../greetings

require (
	ethan.com/greetings v0.0.0-00010101000000-000000000000
	ethan.com/pizzas v0.0.0-00010101000000-000000000000
	ethan.com/time v0.0.0-00010101000000-000000000000
)

replace ethan.com/time => ../time

replace ethan.com/cmd => ../cmd

replace ethan.com/pizzas => ../pizzas
