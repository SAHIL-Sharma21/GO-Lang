//pointers in go lang

package main

import "fmt"

// by value pass ho rha hai num -> pass by value
// func changeNum(num int) {
// 	num = 5 //this is copy of the num passed to it
// 	fmt.Println("In changeNum", num)
// }

// passing by refernece -> to change the original number
// * is the address of that variable
func changeNum(num *int) {
	*num = 5 //de reference the pointer to get the value
	fmt.Println("In Change num functions", *num)
}

func changeSlice(nums *[]int) {
	*nums = append(*nums, 4)
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Println("Memory Address of num:", &num)
	fmt.Println("After changeNum in main", num)

	//slices
	var nums = []int{1, 2, 3}
	changeSlice(&nums)
	fmt.Println("after change", nums)

}
