package main

import "fmt"

func main() {

	days := []string{"sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {

	// 	fmt.Printf("index if  and the day on that day is %v\n", day)

	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		fmt.Println("Value is :", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at ddd")

}
