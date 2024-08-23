//channels -> goroutines ke beech me data bejna hota hai tb channels ke use hota hai
//go routines ke beech mei communication woh channels ki helps se hota hai

package main

import (
	"fmt"
	"time"
)

//sending
// func processNum(numChan chan int) {
// 	for value := range numChan {
// 		fmt.Println("processing num", value)
// 		time.Sleep(time.Second)
// 	}
// }

// for recieve the data
func sum(result chan int, num1 int, num2 int) {
	total := num1 + num2
	result <- total //passing th total to the channel
}

// goroutine synchronizer
func task(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("processing...")
}

// implementing queue sysrtem
func emailSender(emilChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()

	for email := range emilChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second) // 1 second mei process krega
	}
}

func main() {

	// numChan := make(chan int)

	// go processNum(numChan)

	// //sending
	// // numChan <- 6

	// //making infinite loop to send data to channels
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//recieving
	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result //blocking hai yeh
	// fmt.Println(res)

	// time.Sleep(time.Second * 2)

	//creating channel
	// messageChan := make(chan string)

	// //sending data to channel -> data channels ke andert send kr rahe hai
	// messageChan <- "Ping" //channels are blocking

	// //recieve the data
	// mes := <-messageChan
	// fmt.Println("message from the channel", mes)

	//bool channel
	// done := make(chan bool)
	// go task(done)

	// <-done //blocking
	// fmt.Println(isDone)

	//buffered channels -> limited amoutn of data without waiting
	emailChan := make(chan string, 100) //buffered channel
	done := make(chan bool)

	go emailSender(emailChan, done)

	// emailChan <- "sahil@gmail.com"
	// emailChan <- "two@example.com"

	//non blocking
	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	for i := 0; i < 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i) //this is not blocking - bahoutfast send ho jayega
	}
	fmt.Println("done sending email...")

	//closing channel
	close(emailChan) //important to cliose the channel to avoid deadlock error

	<-done //blockingmain goroutine

}
