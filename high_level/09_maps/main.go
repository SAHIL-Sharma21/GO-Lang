package main

import (
	"fmt"
)

// maps or hashtables -> key value pairs
func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["js"] = "Javascript"
	languages["py"] = "Python"
	languages["go"] = "GoLang"

	fmt.Println("maps of languages:", languages)
	fmt.Println("length of maps", len(languages))
	fmt.Println(languages["js"])

	// delete tha value
	// delete(languages, "py")
	// fmt.Println("map after delete", languages)

	//loops in go lang
	for key, value := range languages {
		fmt.Printf("key in language %v, and value is %v\n", key, value)
	}

	//using map literal
	myProfile := map[int]string{
		0: "Hello,",
		2: "this value to be deleetd",
	}
	myProfile[1] = "sahil"

	fmt.Println("myProfile", myProfile)
	// fmt.Println(len(myProfile))

	//delete the value
	delete(myProfile, 2)
	fmt.Println(myProfile)
}
