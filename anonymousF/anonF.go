package main

import "fmt"

// standard or named function
// func add(a, b int) {
// 	fmt.Println(a + b)
// }

func main() {
	// we invoke the add function here
	// add(7, 7)
	// anonymous function
	//Immediately Invoked Function Expression
	// IIFE

	a := 40     //expression
	if a > 20 { //if block
		fmt.Println("40 is grater than 20")
	}

	func(c int, d int) {
		e := c * d
		fmt.Println(e)
	}(7, 9)
}

func init() {
	fmt.Println("I will be executed first")
}
