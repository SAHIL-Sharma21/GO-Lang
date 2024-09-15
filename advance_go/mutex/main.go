package main

import (
	"fmt"

	"net/http"
	"sync"
)

var signals = []string{"test"}

var mut sync.Mutex //pointer of that

var wg sync.WaitGroup

func main() {
	fmt.Println("Mutex in golang")

	websiteList := []string{
		"https://github.com/",
		"https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world",
		"https://google.com",
		"https://instagram.com",
	}

	for _, val := range websiteList {
		go getStatusCode(val)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string) {

	defer wg.Done()

	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)

		mut.Unlock()
		fmt.Printf("%d status code for %s\n", resp.StatusCode, endpoint)
	}

}
