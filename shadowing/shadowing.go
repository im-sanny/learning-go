package main

import "fmt"

var a = 10

func main() {
	age := 60

	if age > 20 {
		a := 300
		fmt.Println(a)
	}
	fmt.Println(a)
}
