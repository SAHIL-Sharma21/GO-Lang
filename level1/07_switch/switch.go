//swithc statement

package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 5

	//in go we can avoid writing break keyword, bhen the scene go added internally.
	switch i {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	default:
		fmt.Println("other")

	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")

	default:
		fmt.Println("Its workDay")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Its a integer")

		case string:
			fmt.Println("Its a string")

		case bool:
			fmt.Println("Its a boolean")

		default:
			fmt.Println("Other")
		}
	}

	whoAmI(28.00)
}
