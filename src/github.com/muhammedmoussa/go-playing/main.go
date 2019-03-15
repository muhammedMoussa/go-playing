package main

import "fmt"

var name = "Muhammed"

func main() {
	// var name = "Muhammed"
	var age int32 = 25
	var isCool = true
	var size float32 = 2.4

	// Shorthand vars
	// name := "Moussa"
	// emial := "me@mail.com"

	name, emial := "Moussa", "me@mail.com"
	fmt.Println(name, emial, isCool, size)
	fmt.Printf("%T\n", age)
}
