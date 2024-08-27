// generics in go lang
package main

import (
	"fmt"
)

// to solve this duplicay we can use generics
func printSlice[T int | string](items []T) { //only int and string can be accpeted
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStringSlic(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type Stack[T int | string] struct {
	elements []T
}

// func (s *Stack[T]) addelement(num int) {
// 	var element T
// 	switch any(element).(type) {
// 	case int:
// 		element = any(num).(T)
// 	case string:
// 		element = any(strconv.Itoa(num)).(T)
// 	}
// 	s.elements = append(s.elements, element)
// }

// comparabele any multiple types
func printAnything[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func main() {

	nums := []int{1, 2, 3, 4}
	names := []string{"Sahil", "Ichigo", "Kempachi"}
	// isAdmin := []bool{true, false, true, true}
	// printStringSlic(names)
	printSlice(names)
	printSlice(nums)
	// printSlice(isAdmin)

	//using on structs
	mystack := Stack[int]{
		elements: []int{1, 3, 5},
	}
	fmt.Println(mystack)

	//string stack
	myStringStack := Stack[string]{
		elements: []string{"Sahil", "Darshana", "Our CAT"},
	}
	fmt.Println(myStringStack)

	printAnything(nums, "Sahil")
	printAnything(names, "Sharma")
}
