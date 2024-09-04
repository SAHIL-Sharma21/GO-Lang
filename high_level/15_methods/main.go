package main

import "fmt"

//functions which in classes is called methods but in go lang we make structs

// making user structs
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// methods
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u *User) NewMail() {
	u.Email = "test@dev.com"
	fmt.Println("Email of this user is", u.Email)
}

// change status -> giving the actual value no the copy
func (u *User) changeStatus() bool {
	u.Status = true
	return u.Status
}
func main() {
	fmt.Println("methods in go lang")

	firstUSer := User{Name: "Sahil", Email: "sahil@dev.com", Status: false, Age: 25}
	fmt.Println("userone  name is", firstUSer.Name)
	fmt.Printf("user details are %+v\n", firstUSer)
	firstUSer.GetStatus()
	firstUSer.NewMail()
	fmt.Println("user email", firstUSer.Email)

	status := firstUSer.changeStatus()
	fmt.Println("new status", status)
	fmt.Println("user status", firstUSer.Status)
}
