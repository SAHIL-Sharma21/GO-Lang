// variadic functions in go lang
package main

import "fmt"

// fnctions which takes n numbers of parameter
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// ...interface{} is like any type in go lang
// func anyFun(data ...interface{}) {

// }

// varidic functions are those in which we can pass n numbers of paramters.
func main() {
	fmt.Println(1, 2, 3, 4, "Sahil") //as Println() recieves the n number of parameter.

	fmt.Println(sum(3, 4, 9))

	result := sum(6, 4)
	fmt.Println(result)

	//passing whole array in the sum funtion
	nums := []int{2, 3, 6}
	answer := sum(nums...)
	fmt.Println(answer)

}
