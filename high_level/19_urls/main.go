package main

import (
	"fmt"
	"net/url"
)

// creating url
const myUrl string = "https://example.com/beginner/baseball?birthday=bottle&bite=beef"

//this is also the type of url
// http://example.com/beginner/baseball?birthday=bottle&bite=beef

func main() {
	fmt.Println("url handling in go lang")
	// fmt.Println(myUrl)
	//parsing url
	result, err := url.Parse(myUrl)

	CheckNilErr(err)
	fmt.Println("scheme of result is:", result.Scheme) //https
	fmt.Println(result.Host)                           //example.com
	fmt.Println(result.Path)                           ///beginner/baseball
	fmt.Println(result.RawQuery)                       //birthday=bottle&bite=beef
	fmt.Println(result.Port())                         //dont have port

	queryParams := result.Query()
	fmt.Printf("The typr of our queryParmas is:  %T\n", queryParams) //url.values -> also known as key value pair

	fmt.Println(queryParams["birthday"]) //[bottle]
	fmt.Println(queryParams["bite"])     //[beef]

	//iterate through key values
	for _, values := range queryParams {
		fmt.Println("Params is", values)
	}

	//passing the reference
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/tutcss",
		RawQuery: "user=sahil",
	}

	//construct the string
	anotherUrl := partsOfUrl.String()
	fmt.Println("another url", anotherUrl) //https://example.com/tutcss?user=sahil

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
