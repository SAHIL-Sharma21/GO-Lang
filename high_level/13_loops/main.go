package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in go lang")

	//making slices to loop over
	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(days)

	//type 1
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println("days are", days[d])
	// }

	//type 2
	for _, value := range days {
		fmt.Println("days", value)
	}
	// for i := range days {
	// 	fmt.Println("Days by:", days[i])
	// }

	myName := "sahil"
	for _, char := range myName {
		fmt.Println("each char is:", string(char))
	}

	//while loop -> implementation
	rougeValue := 1
	for rougeValue <= 10 {

		// if rougeValue == 5 {
		// 	break
		// }

		if rougeValue == 2 {
			goto lco //goto statement
		}

		if rougeValue == 6 {
			rougeValue++ //important case
			continue
		}

		fmt.Println("value is ->", rougeValue)
		rougeValue++
	}

	//goto statement
lco:
	fmt.Println("jumping to lco")

}
