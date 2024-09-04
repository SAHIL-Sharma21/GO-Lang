package main

import "fmt"

func main() {
	fmt.Println("Functions in go lang")
	greeter()

	func() {
		fmt.Println("IIFE in go lang are callled in main func")
	}()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	bigCalc := proAdder(3, 5, 6, 7, 8, 9)
	fmt.Println("big calc is", bigCalc)

	increase, message := proMul(3, 4, 5, 6, 7)
	fmt.Println("increase function gives", increase)
	fmt.Println("message is", message)

}

func greeter() {
	fmt.Println("Hello from greeter")
}

func adder(num1, num2 int) int {
	return num1 + num2
}

// function for many argumnets
// ... are variadic functions
func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// fuintion retruing 2 values
func proMul(values ...int) (int, string) {
	total := 1
	for _, v := range values {
		total *= v
	}
	return total, "This is done by proMul"
}
