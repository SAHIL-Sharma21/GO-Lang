package main

import "fmt"

func main() {
	defer fmt.Println("hello go lang")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Defer in go lang")

	myDefer()
}

//LIFO - last in first out in defer case i.e. we get Defer in go..-> two -> one -> hello go lang

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
