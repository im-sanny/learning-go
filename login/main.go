package main

import (
	"fmt"
)

func welcome() {
	fmt.Println("Welcome!")
}

func userName() string {
	var name string
	fmt.Println("Enter your name -")
	fmt.Scanln(&name)
	return name
}

func userMail() string {
	var mail string
	fmt.Println("Enter your email -")
	fmt.Scanln(&mail)
	return mail
}

func userPassword() string {
	var pass string
	fmt.Println("Enter your password -")
	fmt.Scanln(&pass)
	return pass
}

func getTwoNumber() (int, int) {
	var x int
	var y int
	fmt.Println("Enter first random number -")
	fmt.Scanln(&x)
	fmt.Println("Enter second random number -")
	fmt.Scanln(&y)
	return x, y
}

func add(x int, y int) int {
	sum := x + y
	return sum
}

func main() {
	welcome()
	name := userName()
	mail := userMail()
	pass := userPassword()
	num1, num2 := getTwoNumber()
	total := add(num1, num2)
	fmt.Println(
		"name:", name,
		"email:", mail,
		"password:", pass,
		"total:", total,
	)
}
