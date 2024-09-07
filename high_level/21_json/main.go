package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"courseprice"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              //this - is used for empt field or tthis filed will be not thrown
	Tags     []string `json:"tags,omitempty"` //when there is nil it wont show up
}

func main() {
	fmt.Println("json in go lang")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"Go lang bootcamp", 900, "myWebsite", "abc123", []string{"cloud", "webdev"}},
		{"ReactJs Advance", 600, "company portal", "aBBc103", nil},
		{"Backend Developer(NodeJs)", 500, "YouTube channel", "---", []string{"cloud", "webdev", "Backend", "Distributed system"}},
	}

	//package this data as json data
	jsonData, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}
