package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "tonato"
	fruitList[2] = "cucumber"

	fmt.Println("fruit list is : ", len(fruitList))

	var veglist = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegy list is : ", veglist)
}
