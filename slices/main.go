package main

import (
	"fmt"
)

func main() {
	// fmt.Println("slices")

	// var Fruitlist = []string{"Apple", "tomato", "Peach"}
	// fmt.Printf("type of fruit list %T\n", Fruitlist)

	// Fruitlist = append(Fruitlist, "Mango", "Banana")
	// fmt.Println(" fruit list : ", Fruitlist)

	// // range are always non exclusive
	// Fruitlist = append(Fruitlist[1:3])
	// fmt.Println(Fruitlist)

	// highscore := make([]int, 4)
	// highscore[0] = 23
	// highscore[1] = 123
	// highscore[2] = 263
	// highscore[3] = 983
	// // highscore[4] = 7777;

	// highscore = append(highscore, 555, 666, 777)

	// fmt.Println(highscore)

	// sort.Ints(highscore)
	// fmt.Println(highscore)

	// fmt.Println(sort.IntsAreSorted(highscore))

	// remove a value from slices based on indexs

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
