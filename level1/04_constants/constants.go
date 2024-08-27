// constant in go lang
package main

import "fmt"

// normal variables are declred outside main function
const name = "DK"
const username = "sahil"

func main() {

	//const
	const username string = "Darshana"

	// username = "rahul"

	const userAage = 30

	// fmt.Println(name)
	fmt.Println(username)

	//multiple constant
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(host, port)

}
