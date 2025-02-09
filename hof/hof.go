package main

import "fmt"

/*
rules for higher order function:
any one of the following 3
1. function as parameter
2. function return
3. both
*/

func processOperation(a int, b int, op func(p int, q int)) func(p int, q int) {
	op(a, b)
	return add
}

func call() func(x int, y int) {
	return add
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	processOperation(3, 7, add)

	sum := call()
	sum(8, 9)
}
