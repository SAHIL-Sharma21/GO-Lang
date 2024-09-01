package main

import "fmt"

func main() {
	fmt.Println("Pointers in golang")

	// var myPtr *int
	// fmt.Println("value of ptr is ", myPtr) //nil defualt value

	myNumber := 24

	var ptr = &myNumber                       // referencing means (&)
	fmt.Println("value of actual ptr:", ptr)  //0xc0000120f0
	fmt.Println("value of actual ptr:", *ptr) //24

	*ptr = *ptr + 2
	fmt.Println("new value is", myNumber)

	myname := "sahil"
	fmt.Println("my name is", myname)
	changingName := &myname
	fmt.Println("changing name", *changingName)
	*changingName = "Sharma"
	fmt.Println("new changing name ", *changingName)
	fmt.Println("my name is", myname)
}
