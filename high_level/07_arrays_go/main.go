package main

import "fmt"

func main() {
	fmt.Println("Array in go lang")

	//arr declare

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "peach"

	fmt.Println("fruit list", fruitList)
	fmt.Println("fruit list", len(fruitList))

	var vegList = []string{"aloo", "bhindi", "kaddo", "peas"}
	fmt.Println("veglist is", vegList)
	fmt.Println("length of veglist", len(vegList))
}
