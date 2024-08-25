// mutexes in the go lang
// mutex stands for mutual excultion -> race condition se bachne ke liye
package main

import (
	"fmt"
	"sync"
)

// views race condition mei aa gya hai ek hi resource ko goroutine update kr raha hai
type post struct {
	views int
	mu    sync.Mutex //making mutex here
}

// views ko increment krn ahia
func (p *post) increment(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock() //locking the variable and other go routine will wait utnil the current modification done!!!
	p.views += 1
}

func main() {

	var wg sync.WaitGroup //making waitgroup

	myPost := post{views: 0}
	// myPost.increment()
	// myPost.increment()

	//if multiple baar ho raha hai -> handling concurrently
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.increment(&wg)
	}

	wg.Wait() //wait for all the goroutines

	fmt.Println(myPost.views)

}
