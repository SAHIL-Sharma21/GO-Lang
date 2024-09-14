package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usually these are pointers

func main() {
	fmt.Println("Waitgroups in go lang")

	websiteList := []string{
		"https://github.com/",
		"https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world",
		"https://google.com",
		"https://instagram.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() //not allow the main method to finish

}

func getStatusCode(endpoint string) {
	defer wg.Done() //passing signal when the work is done

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS... in endpoint")
	}

	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
