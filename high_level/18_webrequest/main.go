package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("web request in go lang")

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("type of response is: %T\n", response)

	defer response.Body.Close() //calller's responsibility to close the connection

	if response.StatusCode == 200 {
		data, err := io.ReadAll(response.Body)
		checkNilErr(err)
		fmt.Println("response from the web requets is: ", string(data))
	}
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
