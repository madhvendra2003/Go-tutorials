package main

import "fmt"

func main() {

	var username string = "Madhvendra Singh"
	var IsloggedIN bool = false
	fmt.Println(IsloggedIN)
	fmt.Printf("Variable is of type : %T \n", username)

	var smallVal int = -455
	var yesNo bool = false
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", yesNo)

	var smallfloat float32 = -455.55555

	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type : %T \n", smallfloat)

	var another int
	fmt.Println(another)

	var stt string
	fmt.Println(stt)

	//implicit way

	var website = "Learncodeonline.in"
	fmt.Printf("Variable is of type : %T \n", website)

	number := 43.0
	fmt.Println(number)

}
