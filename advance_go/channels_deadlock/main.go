package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and deadlock in go")

	//creating a channel
	myChan := make(chan int, 2) //buffered channel
	wg := &sync.WaitGroup{}
	//making error of deadlock
	// myChan <- 2
	// fmt.Println("my chan", <-myChan)

	wg.Add(2)
	//solution -> making goroutines
	//adding the values
	//recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// fmt.Println("value from the channel:", <-ch)
		val, isChannelOpen := <-ch
		val2 := <-ch

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		fmt.Println(val2)

		// fmt.Println(<-ch)
		wg.Done()
	}(myChan, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		ch <- 4
		ch <- 5
		close(ch) //close channel
		// ch <- 5
		// ch <- 6
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
