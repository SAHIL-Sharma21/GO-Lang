// maps in go lang
package main

import (
	"fmt"
	"maps"
)

// maps -> Hashmap, Objects, Dict
func main() {
	//creating map
	//map[key_Type]Value_Type

	programmingLang := make(map[string]string)

	//setting an element
	programmingLang["one"] = "Go Lang"
	programmingLang["two"] = "Rust"
	fmt.Println(programmingLang)

	//get an element
	fmt.Println(programmingLang["one"]) //Go Lang

	//IMP: if key does not exist in map then it retuins zero value (if string -> "", numer -> 0, bool -> false)
	fmt.Println(programmingLang["three"]) // we get emoty value

	projects := make(map[int]string)
	projects[1] = "Frontend"
	projects[2] = "Backend"
	projects[3] = "FullStack"
	projects[4] = "AI"

	fmt.Println(projects)
	fmt.Println(projects[2]) // backend

	//map length
	fmt.Println(len(projects)) // 2

	//delete the element from the map
	delete(projects, 3)
	fmt.Println(projects) //map[1:Frontend 2:Backend 4:AI]

	//empty full map
	// clear(programmingLang)
	// fmt.Println(programmingLang) // map[]

	//another way to make map
	m := map[string]int{"price": 30, "discount": 20, "userCount": 5}
	fmt.Println(m) //map[discount:20 price:30 userCount:5]

	//checking the element in the map
	// _, ok := m["price"]
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("Not Ok")
	// }

	value, isPresent := m["userCount"]
	if isPresent {
		fmt.Printf("the value is %d \n", value)
	} else {
		fmt.Println("Element not found")
	}

	//comparing two maps
	m1 := map[int]string{1: "Sahil", 2: "Avinash"}
	m2 := map[int]string{1: "Sahil", 2: "Avinash"}

	//maps is inbuilt
	fmt.Println(maps.Equal(m1, m2)) //true
	m3 := map[int]string{}
	maps.Copy(m3, m1)
	fmt.Println(m1)
	fmt.Println(m3)

}
