package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(5, 8)
	add(p, a)
}

func main() {
	call()

	fmt.Println(a)
}

func init() {
	fmt.Println("Hello!")
}

/*
2 phase:
 1. compilation
 2. execution


	********* Compile Phase *********

	** code segment **
	a = 10
	call = func (){...}
	add = func (){...}
	main = func (){...}
	init = func (){...}



 go run main.go => compile it => main => ./main
 go build main.go => compile it => /main

 ./main
*/
