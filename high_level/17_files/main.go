package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File handling in golang")

	content := "This needs to go line file"

	file, err := os.Create("./myfiles.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("Length is", length)

	defer file.Close()

	//reading the file here
	readFile(file.Name())
}

// reading the file
func readFile(filaname string) {
	content, err := os.ReadFile(filaname)
	checkNilErr(err)
	fmt.Println("file content are:", string(content))
}

// best way to handle error:
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
