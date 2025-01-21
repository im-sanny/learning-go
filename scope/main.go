package main

import "fmt"

var a = 10
var b = 20

func main() {
	x := 18

	if x >= 18 {
		p := 1
		fmt.Println("I am young enough to get married")
		fmt.Println("I have", p, "girlfriend")
	}

	age := 18
	
	switch {
	case age >= 8:
		fmt.Println("u should watch toy story")
	default:
		fmt.Println("go to play")
	}
}
