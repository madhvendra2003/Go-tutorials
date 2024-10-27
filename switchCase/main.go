package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch Case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Vlaue of the dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 6:
		fmt.Println("Congratulation you have gotten 6 , move 6 spots and take another chance")

	default:
		if diceNumber <= 5 {
			fmt.Printf("move %v spaces \n", diceNumber)
		} else {
			fmt.Println("please take another chance")
		}
	}

}
