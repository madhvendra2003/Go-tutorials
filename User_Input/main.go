package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Madhvendra singh"
	fmt.Printf(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading for our pizza : ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)
	fmt.Printf("type of rating, %T", input)

}
