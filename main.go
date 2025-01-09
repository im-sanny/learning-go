package main

import "fmt"

func main() {
	age := 18
	sex := "male"

	if age == 20 && sex == "male" {
		fmt.Println("You're ready for war")
	} else if age >= 20 && sex != "male" {
		fmt.Println("Sorry we're looking for male soldiers")
	} else {
		fmt.Println("Try again later")
	}
	main1()
}

func main1() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("either 3 or more")
	}
}
