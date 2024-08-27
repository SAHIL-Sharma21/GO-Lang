//range in go lang

package main

import "fmt"

// range is used for iterating over Data structure
func main() {
	// nums := []int{1, 2, 4, 6, 8}

	//using for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// //using range keyword
	// for _, num := range nums {
	// 	sum += num
	// 	// fmt.Println(num)
	// }
	// fmt.Println(sum)

	// for i, num := range nums {
	// 	// fmt.Printf("the index of the arr is %d and value is %d. \n", i, num)
	// }

	// //range iterations in maps
	// map1 := make(map[string]int)
	// map1["name"] = 1
	// fmt.Println(map1)
	// fmt.Println(len(map1))

	m := map[string]string{"fname": "John", "lname": "Doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//to get only key
	for key := range m {
		fmt.Println(key)
	}

	//range in string -> unicode code point rune
	// here i is the starting byte of the rune
	//255 -> 1 byte --- if num is greaster then 255 then it will take 2 byte
	for i, chr := range "Sahil" {
		fmt.Printf("the index of char is %d and value is %d\n", i, chr)
	}

	for _, char := range "Darshana" {
		fmt.Println(string(char))
	}

}
