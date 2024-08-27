// enums in go lang
package main

import "fmt"

// enums are enumerated types
// in go lang there is no enums but we can implemmet by const keyword and make custom types
type OrderStatus int

const (
	Received OrderStatus = iota //iota is zero based index initially zero and further increases automatically
	Confirmed
	Prepared
	Delivered
)

// using string
type SelectSeat string

const (
	Window SelectSeat = "window"
	Side              = "side"
	Aisle             = "aisle"
)

// example
func changeStatus(status OrderStatus) {
	fmt.Println("Changing status to", status)
}

func getSeat(seatChoice string) {
	fmt.Println("You are assigned with seat", seatChoice)
}

func main() {
	changeStatus(Prepared)
	getSeat(Aisle)
}
