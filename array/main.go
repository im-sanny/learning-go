package main

import "fmt"

var arr2 =[3]string {"I", "Love", "You"}

func main() {
	var arr [2]int

	arr[1] = 6
	arr[0] = 7

	fmt.Println(arr)
	fmt.Println(arr2[0])
}
