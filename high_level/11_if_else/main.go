package main

import "fmt"

func main() {
	fmt.Println("If else in go lang")

	loginCount := 26
	var message string

	if loginCount < 10 {
		message = "Regular user"
	} else if loginCount > 10 {
		message = "not a normal user"
	} else {
		message = "hackeer is there!!"
	}

	fmt.Println("message:", message)

	//another way
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	//special syntax -> assiging on the go
	if num := 4; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
