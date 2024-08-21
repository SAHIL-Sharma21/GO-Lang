//structs in go lang
//structs is a custom data structure in go lang its same as classes in other language

package main

import (
	"fmt"
	"time"
)

// order structs
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nano seconds precision of this time
}

// for making constructor we use hack and make one function
func newOrder(id string, amount float32, status string) *order {
	//initial setup here
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder //we are returing poiter of that struct
}

// attaching the method in the structs -> that bracket is reciever type which we have go give as it is convention to write the first letter of the structs followed by name of that struct
// reciever type
func (o *order) changeStatus(status string) {
	o.status = status // structs automatically De-referencing ka kaam kr deta hai
}

func (o order) getAmountPrice() float32 { //removed ptr * inncase of getting the value
	return o.amount
}

func main() {

	//making instance
	// userOrder := order{
	// 	id:     "1",
	// 	amount: 45.00,
	// 	status: "recieved",
	// }

	// userOrder.createdAt = time.Now()

	// fmt.Println("order struct:", userOrder)

	// //getting single field
	// fmt.Println(userOrder.amount)

	// //more instance of the struct -> they are different insatnce and does not affect other instance
	myOrder := order{
		id:        "2",
		amount:    90,
		status:    "Dispatched",
		createdAt: time.Now(),
	}

	// fmt.Println("second order:", myOrder)
	myOrder.changeStatus("Delivered")
	fmt.Println(myOrder)

	fmt.Println(myOrder.getAmountPrice()) //90

	//if we dont set any value to the field then, Default values is zero value
	//zero values are : int => 0, float => 0, string => "", bool => false
	// ecomOrder := order{
	// 	id: "3",
	// 	// amount: 20,
	// 	// status: "",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println(ecomOrder)

	//using contructor way to make instance
	userOrder := newOrder("4", 35, "cancelled")
	fmt.Println(userOrder)

	//inline structs if we have to use only for 1 time use
	//making new structs
	language := struct {
		name   string
		isGood bool
	}{"Go Lang", true}
	fmt.Println(language)

}
