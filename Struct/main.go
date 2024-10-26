package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; no super or parent

	hitesh := User{"Madhvendra", "1789@gmail.com", true, 13}
	fmt.Println(hitesh)

	fmt.Printf("Madhvendra details are : %+v\n ", hitesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
