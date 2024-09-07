package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET request in go lang")
	// PerformGetRequest()
	PerformPostJsonRequest()
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

// POST request
func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/getcountry"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Go lang course",
			"price": 0,
			"Demand":"High"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("data", string(content))
}
