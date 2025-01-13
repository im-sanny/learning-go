package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func printSomething() {
	fmt.Println("Education must be free!")
}

func greetings(name string) {
	fmt.Println("Hello Habib vai, thanks for the course,", name)
}

func main() {
	printSomething()
	greetings("Sanny")
}
