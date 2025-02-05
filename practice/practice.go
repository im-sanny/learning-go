package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to the application")
}

func getUserName() string {
	var name string
	fmt.Println("Please enter ur name -")
	fmt.Scanln(&name)
	return name
}

func getTwoNumber() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter the first number -")
	fmt.Scanln(&num1) //& => ampersand
	fmt.Println("Enter the second number -")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func display(name string, sum int) {
	fmt.Println("Hello,", name)
	fmt.Println("Summation =", sum)
}

func farewellMessage() {
	fmt.Println("Thanks for using the application!")
	fmt.Println("Goodbye!")
}

func main() {
	// Print welcome message
	welcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumber()
	sum := add(num1, num2)
	display(name, sum)
	farewellMessage()
}
