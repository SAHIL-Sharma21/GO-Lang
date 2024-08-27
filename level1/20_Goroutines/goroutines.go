// goroutines in go lang -> one of the powerful feature of go lang

// goroutines are the lightweight threads which allows us to do multithreading -> to run concurrently
package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task of id", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		// task(i) //normal call
		//goroutine call -> power of golang
		// go task(i)

		go func(i int) {
			fmt.Println(i)
		}(i)

	}

	//to simulate and see the result we will sleep the main func for some time as main func is also run in goroutine
	time.Sleep(time.Second * 2)
}
