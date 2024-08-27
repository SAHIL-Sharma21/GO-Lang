//arrays

package main

import "fmt"

func main() {

	//zeroed values will be there
	//int -> 0, string -> "", bool -> false
	var nums [4]int

	//array length
	println(len(nums)) //4

	nums[0] = 1

	fmt.Println(nums[0]) //1
	fmt.Println(nums)    // //array length
	println(len(nums))   //[1 0 0 0]

	//in case of bool array it get all fasle
	var values [4]bool
	values[2] = true
	fmt.Println(values) // [false false false false]

	//strings array
	var usernames [3]string
	usernames[0] = "Darshana"
	usernames[1] = "Sahil"
	fmt.Println(usernames) //[Darshana Sahil ]

	//array and giving values when declaring
	//to declared it in single line
	numbered := [3]int{2, 4, 6}
	fmt.Println(numbered) //[2 4 6]

	//2D arrays
	chess := [2][2]int{{2, 4}, {6, 8}}
	fmt.Println(chess) //[[2 4] [6 8]]

	// - fixed size, then use arrays, that is predictazble
	// - memeory optimization can be enhanced
	// - constant time access O(1)

}
