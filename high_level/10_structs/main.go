package main

import "fmt"

func main() {
	fmt.Println("structs in go lang")
	//no inhertiance in go lang and there is no classes in go lang; no Super or parent
	sahil := User{
		Name:   "Sahil",
		Email:  "sahil@dev.com",
		Status: true,
		Age:    25,
	}

	fmt.Println("sahil by struct", sahil)
	fmt.Printf("sahil details are: %+v\n", sahil) //sahil details are: {Name:Sahil Email:sahil@dev.com Status:true Age:25}
	fmt.Printf("Name is %v and email is %v\n", sahil.Name, sahil.Email)
}

// structs declartion
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
