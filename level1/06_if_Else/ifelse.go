//if esle in go

package main

import "fmt"

func main() {
	userAge := 16

	if userAge >= 18 {
		fmt.Println("person is an adult")
	} else if userAge >= 12 {
		fmt.Println("person is teenage")
	} else {
		fmt.Println("pserson is not an adult")
	}

	userRole := "admin"
	userPermission := true

	if userRole == "admin" || userPermission {
		fmt.Println("user is admin")
	}

	//we can declare a variable in inside if construct

	if age := 15; age >= 18 {
		fmt.Println("person is adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenage", age)
	} else {
		fmt.Println("person is kid", age)
	}

	//go deos not have ternary operator, we will have to use nornmal if else.
}
