package main

import (
	"bufio"
	"fmt"

	"os"

	"github.com/fatih/color"
)

func main() {
	welcome := "Welcome user input"
	color.Blue(welcome)

	//user input can be done through bufio and os pkg
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our app:")

	//comma OK or err, OK syntax in go lang
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T\n", input)
}
