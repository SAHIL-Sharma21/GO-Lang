package main

import "fmt"

func changeName(newName string) {
	newName = "Ichigo"
}

func changeUserName(newName *string) {
	fmt.Println(*newName) //sahil
	*newName = "Ichigo"
}

func main() {

	// value := 35

	// fmt.Println("The value is", value)
	// fmt.Println("The Address is", &value)

	name := "sahil"
	fmt.Println(name)
	//passing by reference
	// changeName(name)
	// fmt.Println(name)

	//passing reference -> pointer
	changeUserName(&name)
	fmt.Println(name)
}
