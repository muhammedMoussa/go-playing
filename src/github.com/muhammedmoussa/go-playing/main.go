package main

import (
	"fmt"
)

var name = "Muhammed"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	// var name = "Muhammed"
	// var age int32 = 25
	// var isCool = true
	// var size float32 = 2.4

	// Shorthand vars
	// name := "Moussa"
	// emial := "me@mail.com"

	// name, emial := "Moussa", "me@mail.com"
	// fmt.Println(name, emial, isCool, size)
	// fmt.Println(math.Sqrt(2))
	// fmt.Printf("%T\n", age)
	fmt.Println(greeting("muhammed"))
	fmt.Println(getSum(70, 4))

}
