package main

import (
	"fmt"

	"github.com/derailed/nuclio/movie"
)

func main() {
	fmt.Println(movie.LoadMem())
	// fmt.Println("Hello World Emoji!")

	// emoji.Println(":beer: Beer!!!")

	// pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
	// fmt.Println(pizzaMessage)
}
