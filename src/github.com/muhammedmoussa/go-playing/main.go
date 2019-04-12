package main

import "fmt"

// var name = "Muhammed"

// // func greeting(name string) string {
// // 	return "Hello " + name
// // }

// // func getSum(num1, num2 int) int {
// // 	return num1 + num2
// // }

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

	// fmt.Println(greeting("muhammed"))
	// fmt.Println(getSum(70, 4))

	// Arrays
	// var fruitArray [2]string
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"
	// nameArray := [2]string{"Muhammed", "Moussa"}

	// langSlice := []string{"Js", "C++", "Python"}

	// fmt.Println(fruitArray, nameArray)
	// fmt.Println(langSlice)

	// Conditions
	// x := 50
	// y := 10
	// color := "blue"

	// if x < y {
	// 	fmt.Printf("%d is less than %d\n", x, y)
	// } else {
	// 	fmt.Printf("%d is less than %d\n", y, x)
	// }

	// if color == "red" {
	// 	fmt.Println("Color is red")
	// } else if color == "blue" {
	// 	fmt.Println("Color is blue")
	// } else {
	// 	fmt.Println("Color undefined")
	// }

	// switch color {
	// case "red":
	// 	fmt.Println("Color is Red")
	// case "blue":
	// 	fmt.Println("Color is Blue")
	// default:
	// 	fmt.Println("Color is undefined")
	// }

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("num %d\n", i)
	// }

	/* MAPS */
	// emails := make(map[string]string)
	// emails["muhammed"] = "muhammed@gmail.com"
	// emails["ahmed"] = "ahmed@gmail.com"
	emails := map[string]string{"muhammed": "muhammed@gmail.com", "ahmed": "ahmed@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["muhammed"])
	fmt.Println(len(emails))

	delete(emails, "ahmed")
	fmt.Println(len(emails))
}
