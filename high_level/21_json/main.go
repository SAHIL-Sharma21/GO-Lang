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
	// EncodeJson()
	DecodeJson()
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

// decode json
func DecodeJson() {
	jsonFromWeb := []byte(`
		{
                "coursename": "Go lang bootcamp",
                "courseprice": 900,
                "website": "myWebsite",
                "tags": ["cloud","webdev"]
        }
	`)

	//to check correct json data
	var proCourse course

	checkValid := json.Valid(jsonFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonFromWeb, &proCourse)

		fmt.Printf("%#v\n", proCourse)
	} else {
		fmt.Println("JSON was not VALID!")
	}

	//some case where we want to add data as key value pair

	//interface{} is a method signature and any value can be there for the key in below map
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &myOnlineData)
	// fmt.Printf("%#v\n", myOnlineData)

	//for key value use loop
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and TYpe is %T\n", k, v, v)
	}

}
