//wait group in go lang

package main

import (
	"fmt"
	"sync"
)

//wait group step is -> add, done and wait

func runTask(num int, w *sync.WaitGroup) {
	defer w.Done() // function execute hone ke baad last mei defer hota hai -> like a cleaning function
	fmt.Println("Doing Task", num)
}

func main() {

	//creating wait group
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go runTask(i, &wg)
	}

	wg.Wait()
}
