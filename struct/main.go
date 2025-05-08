package main

import "fmt"

type User struct {
	Name string //member variable or property
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

func (usr User) printDetails() {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

func (usr User) call(a int) {
	fmt.Println("Name:", usr.Name)
	fmt.Println(a)
}

//printUserDetails(user1)
//user1.printDetails()

func main() {
	var user1 User

	user1 = User{
		Name: "Habib",
		Age:  30,
	}

	user1.printDetails()
	user1.call(1)

	user2 := User{ //instance or object
		Name: "Bingo",
		Age:  20,
	}

	printUserDetails(user2)
}

/*
2 Phase
	1. Compilation phase (compile time)
	2. execution phase (run time)

	********** Compile Phase ***********

	** Code Segment **

	User = type User Struct{...}
	printUserDetails = func (usr User){...}
	printDetails = func(){...}	//User
	call = func(a int){...}	//User
	main = func (){...}


 go run main.go => compile it => main => ./main
 go build main.go => compile it => /main

 ./main
*/
