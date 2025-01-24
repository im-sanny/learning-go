package main

import "fmt"

var a = 10

func main() {
	age := 30

	if age > 20 {
		a := 50
		fmt.Println(a)
	}
	fmt.Println(a)
}
