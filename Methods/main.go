package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	hitesh := User{"Madhvendra", "1789@gmail.com", true, 13, 5}

	hitesh.email = "mas@123gmail.com"

	fmt.Printf("Madhvendra details are : %+v\n ", hitesh)
	hitesh.GetStatus()

}

type User struct {
	Name   string
	email  string
	Status bool
	Age    int
	OneAge int
}

func (u User) GetStatus() {
	fmt.Println(u.OneAge)
}

func (u User) getMail() string {

	return u.email

}
