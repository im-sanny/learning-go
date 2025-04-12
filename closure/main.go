package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("==Bank==")
}

/*
2 phase:
 1. compilation (compile time)
 2. execution (run time)


	********* Compile Phase *********

	** code segment **
	a = 10
	outer = func() {...}
	outerAnonymous = func() {...}
	call = func() {}
	main = func(){}
	init = func(){}




 go run main.go => compile it => main => ./main
 go build main.go => compile it => /main

 ./main
*/
