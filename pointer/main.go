package main

import "fmt"

func main() {
	//pointer or address of memory (ram)

 	x := 20

	p := &x // ampersand & => address of |

	fmt.Println("address:", p) 								//p is the address of x
	fmt.Println("value at the address:", *p) //value at the address
}
