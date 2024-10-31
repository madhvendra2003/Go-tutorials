package main

import "fmt"

func main() {
	greet()
	greeter2()

	fmt.Println("the result is ", sum(2, 3))

	fmt.Println("the new result is", proAdder(1, 2, 3, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 4, 2233, 23, 2323, 2))

}

func greeter2() {
	fmt.Println("hello this is me too")
}

func greet() {
	fmt.Println("hello this is me")
}

func sum(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}
