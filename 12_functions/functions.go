//function in go lang

package main

import "fmt"

// making functio to add integers
func add(num1 int, num2 int) int {
	return num1 + num2
}

// anotehr way -> if the paramater type is same then we can give only 1 int at the end and uske phele jitne bhi paramater int type ke honge
func sub(num1, num2 int) int {
	return num2 - num1
}

// go can return mutiple valuies
func getLanguages() (string, string, int) {
	return "GoLang", "Rust", 10
}

// function recivuing another function and retuiting another function
// func processIt(fn func(a int) int) {
// 	fn(1)
// }

// function retun another funtion
func processIt() func(a int) int {
	return func(a int) int {
		return 5
	}
}

func main() {
	v1 := 4
	v2 := 10

	fmt.Println(add(v1, v2))
	result := sub(v1, v2)
	fmt.Println(result)

	// fmt.Println(getLanguages())
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3) // GoLang Rust 10

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	//process it will retun funtion
	getValue := processIt()
	fmt.Println(getValue(6))

}
