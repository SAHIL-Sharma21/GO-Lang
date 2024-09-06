package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET request in go lang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/healthcheck"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //alwyas do this

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("content length", response.ContentLength)

	//another way we can create a builder from strings std pkg
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println("byte count: ", byteCount) //41

	fmt.Println("data", responseString.String()) //data {"message":"server is up and running..."}
	// fmt.Println(string(content))
}
