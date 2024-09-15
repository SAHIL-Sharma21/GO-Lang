package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in go")

	wg := &sync.WaitGroup{} //good practice

	//to solve the issue we need mutex to do it
	mut := &sync.Mutex{} //used whenwe are writing some data

	var score = []int{0}

	wg.Add(3) //how many goroutines are there we will pass that value
	// wg.Add(1)
	//creating go routine -> IIFE
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("1st go routine")
		mut.Lock()               //lock the mem
		score = append(score, 1) //write operation
		mut.Unlock()             // unlock the mem
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("2nd go routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("2nd go routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
