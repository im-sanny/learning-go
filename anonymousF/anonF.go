package main

import "fmt"

// standard or named function
// func add(a, b int) {
// 	fmt.Println(a + b)
// }

func sum() {
	add(3, 5)
}
func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	// we invoke the add function here
	// add(7, 7)
	// anonymous function
	//Immediately Invoked Function Expression
	// IIFE

	// a := 40     //expression
	// if a > 20 { //if block
	// 	fmt.Println("40 is grater than 20")
	// }
	
	sum()

	func(c int, d int) {
		e := c * d
		fmt.Println(e)
	}(7, 9)

	// function expression or assign function in variable
	plus := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}
	plus(3, 5)
}

func init() {
	fmt.Println("I will be executed first")
}
