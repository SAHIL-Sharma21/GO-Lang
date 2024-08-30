package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths")

	// var mynumberone int = 2
	// var mynumbertwo float64 = 25.6
	// fmt.Println("The sum is", mynumberone+int(mynumbertwo)) //this is bad practice

	//random numbers -> by maths/rand
	// rand.Seed(time.Now().UnixNano())// this rand.seed is deprecated
	// rand.NewSource(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random value from crypto
	myRandomInt, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		panic(err)
	}
	fmt.Println("Random vlaue", myRandomInt)

}
