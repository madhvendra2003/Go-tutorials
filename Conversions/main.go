package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza shop")
	fmt.Println("please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating, ", input)

	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to yout rating : ", numRating+1)
	}

}
