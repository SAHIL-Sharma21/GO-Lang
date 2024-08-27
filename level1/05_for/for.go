// for loop
package main

import "fmt"

//go only have for loop for iterations, there is not while loop

// for -> only construct in for go  for looping
func main() {
	//implementing while loop in go with for loop

	//while loop
	// i := 0
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	println(1)
	// }

	//classic for loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	//break and continue
	for i := 0; i <= 5; i++ {
		if i == 4 {
			continue
		}

		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//go new feature has range
	for i := range 4 {
		fmt.Println(i)
	}
}
