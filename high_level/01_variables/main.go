package main

import "fmt"

// token := 4774 -> we cannot use valorous operator here it can be only used inside the funtions.

// constant
const LoginToken string = "idsjdjsdkslkljalkldck" // -> this is puclic variable in go this is the way to declare

func main() {

	var username string = "Sahil"
	fmt.Println(username)
	fmt.Printf("variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("isLoggedIn", isLoggedIn)
	fmt.Printf("variable is of type %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variable is of type %T \n", smallValue)

	var smallFloatVal float32 = 255.3414556456435643
	fmt.Println(smallFloatVal) //255.34146
	fmt.Printf("variable is of type %T \n", smallFloatVal)

	var largeFloatVal float64 = 255.3414556456435643
	fmt.Println(largeFloatVal) //255.34145564564358
	fmt.Printf("variable is of type %T \n", largeFloatVal)

	//default values and aliases
	var newvar int
	fmt.Println(newvar)                              // 0
	fmt.Printf("variable is of type: %T \n", newvar) //int

	//implicit way of declaring a variable
	var myname = "sahil"
	fmt.Println(myname)

	//no var style
	// := is call valarous operator
	numberOfUser := 3000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
