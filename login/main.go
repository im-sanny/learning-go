package main

import "fmt"

func welcome() {
	fmt.Println("Welcome!")
}

func userName() string {
	var name string
	fmt.Println("Enter your name -")
	fmt.Scanln(&name)
	return name
}

func userMail() string {
	var mail string
	fmt.Println("Enter your email -")
	fmt.Scanln(&mail)
	return mail
}

func userPassword() string {
	var pass string
	fmt.Println("Enter your password -")
	fmt.Scanln(&pass)
	return pass
}



func main() {
	welcome()
	userName()
	userMail()
	userPassword()

}
