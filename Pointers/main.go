package main

import "fmt"

func main() {
	fmt.Println("Madhvendra Singh")

	// var ptr *int

	// fmt.Println(ptr)

	myNumber := 56

	var ptr *int = &myNumber

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println(myNumber)
}
