package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("goroutines in go lang -> advance topic")

	go greeter("hello")
	go greeter("Golang")

	time.Sleep(2 * time.Millisecond)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
