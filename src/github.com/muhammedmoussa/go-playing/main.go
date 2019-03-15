package main

import (
	"fmt"
	"math"

	"github.com/muhammedmoussa/go-playing/strutil"
)

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
	fmt.Println(math.Sqrt(2))
	fmt.Printf("%T\n", age)
	fmt.Println(strutil.Reverse("olleH"))
}
