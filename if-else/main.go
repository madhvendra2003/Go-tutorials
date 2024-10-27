package main

import "fmt"

func main() {

	fmt.Println("Madhvendra Singh")

	loginCount := 11

	var temp string

	if loginCount < 10 {
		temp = "I can handle it in my sleep"
	} else if loginCount < 20 {
		temp = "this is no hard"
	} else {
		temp = "this is too much"
	}

	result := temp

	fmt.Println(result)

}
