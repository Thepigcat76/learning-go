package main

import (
	"fmt"
	"learning_go/lib"
)

func main() {
	fmt.Println("hi")
	lib.Input("Your age: ", "An error occured while reading the input, please try again", "Your age is: ")
}