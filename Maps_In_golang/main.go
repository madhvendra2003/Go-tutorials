package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["Py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("js is short for : ", languages["js"])

	delete(languages, "rb")

	fmt.Println("List of all languages: ", languages)
	fmt.Println("js is short for : ", languages["js"])

	// loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
