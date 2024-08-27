//varibales in go Lang

package main

import "fmt"

func main() {

	//string var
	var username string = "Sahil" //we have to use this variable in go lang

	//go lang can infer type automatically
	var userAge = 25

	var isAdmin = true

	// fmt.Println(username)
	fmt.Println(userAge)

	if isAdmin {
		fmt.Println(username)
	}

	//shorthand syntax

	name := "go lang"
	fmt.Println(name)

	//use case -> definig the type and then assigning after some use case
	var userPref string

	userPref = "normal"
	fmt.Println(userPref)

	//floats
	// var price float32 = 25.90

	price := 28.90
	fmt.Println(price)
}
