package main

import (
	"fmt"
	"slices"
)

//slices - dyanamic array
//most used construct - always use in go lang
// we get

func main() {

	//declaring a slices -> uninitilized slice in nil in go lang (nil means null)
	// var nums []int
	// fmt.Println(nums)        // []
	// fmt.Println(len(nums))   // 0
	// fmt.Println(nums == nil) //true

	//2nd way -> 2 is the initial size of the slice
	var numbers = make([]int, 2)

	fmt.Println(numbers)        // [0 0]
	fmt.Println(numbers == nil) // false
	fmt.Println(cap(numbers))   //2 cap is capacity same as len
	//capacity-> maximum number of elementcan fit meaning

	var nums1 = make([]int, 2, 5)
	fmt.Println(cap(nums1)) // 5 capacity but automaticaaly resizes it when we add element

	nums1 = append(nums1, 3, 4, 5, 6, 7, 8)
	fmt.Println(nums1)      //[0 0 2]
	fmt.Println(cap(nums1)) //5

	//corrct way of doing things
	var price = make([]int, 0, 5)
	price = append(price, 2)
	price = append(price, 4, 6)
	fmt.Println(price)      // [2 4]
	fmt.Println(len(price)) // 2

	//different way to make slices
	userName := []string{}
	fmt.Println(userName)      //[]
	fmt.Println(cap(userName)) //0
	fmt.Println(len(userName)) // 0

	userName = append(userName, "sahil", "Darshana")
	fmt.Println(userName)
	fmt.Println(len(userName), cap(userName)) // 2 2

	userName[1] = "Malak"
	fmt.Println(userName)

	var nums = make([]int, 0, 5)
	nums = append(nums, 2)
	nums = append(nums, 4)
	var nums2 = make([]int, len(nums))

	//copy function
	fmt.Println(nums, nums2) //[2 4][0 0]

	//copy func
	copy(nums2, nums)
	fmt.Println(nums, nums2) //[2 4][2 4]

	//slice operator
	var nums3 = []int{1, 2, 3}
	fmt.Println(nums3[0:2]) //[1 2]
	fmt.Println(nums3[:1])  //[1]
	fmt.Println(nums3[2:])  //[3]

	//slice inbuilt package
	var nums4 = []int{1, 2}
	var nums5 = []int{1, 2}

	fmt.Println(slices.Equal(nums4, nums5)) //true

	//2D array
	var nums6 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums6)      //[[1 2 3] [4 5 6]]
	fmt.Println(len(nums6)) // 2
	fmt.Println(nums6[1])
}
