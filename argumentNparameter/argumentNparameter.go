package main

import "fmt"

func add(a int, b int) { // parameter => a, b
	c := a + b
	fmt.Println(c)
}

func main() {
	add(6, 8) //argument => 6, 8
}

/*
1. parameter vs argument
2. first oder function
	i. standard or named function
	ii. anonymous function
	iii. IIFE
	iv.function expression
3. Higher order function or first class function
4. callback function
5. first class citizen -> (data assigned in variable)

functional paradigm -> haskell, racket

math -> logic(discrete mathematics)

1. first order logic
2. higher order logic

### logic ###
1. Object (people, animal, car etc)
2. Property (color, student)
3. Relation

Mutul is a student
Apple is red
Mutul is taller than Kutul

Statement

### First order logic

Rules: All customer must pay their pizza bills
			 All students must wear their uniforms.

### Higher order logic
Any rules that applies on all customer must apply on Mutul also
a rule: customer can't give tips in this restaurant
*/
