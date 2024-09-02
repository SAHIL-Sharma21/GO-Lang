package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in go lang")
	//mostly used data type im go lang

	//slices in go lang
	var fruitList = []string{"Apple", "Cheery", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Avacado")

	// fmt.Println("Fruit list", fruitList)
	// fmt.Println("length of the fruit list is", len(fruitList))

	// fruitList = append(fruitList[1:])
	// fruitList = append(fruitList[:2])
	// fmt.Println(fruitList)

	//using new and make keyword
	highScores := make([]int, 5)
	highScores[0] = 2
	highScores[2] = 48
	highScores[4] = 100
	// highScores[5] = 200 //-> index out of bound err

	//but we can append in the slice
	highScores = append(highScores, 150, 200)
	fmt.Println("high score slice", highScores)
	fmt.Println("high score length", len(highScores))

	sort.Ints(highScores) //sort the slice in aesc
	fmt.Println("sorted slice", highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //true

	//how to remove a value from slice based on index
	var courses = []string{"Golang", "Rust", "Swift", "Dart", "Zig"}
	fmt.Println(courses)

	index := 2                                              //to be removed
	courses = append(courses[:index], courses[index+1:]...) //removing using append method as slicing the slice
	fmt.Println("courses", courses)

}
