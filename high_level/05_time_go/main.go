package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in go lang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-10-2001 15:00:01 Monday"))

	createdDate := time.Date(2024, time.August, 14, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
