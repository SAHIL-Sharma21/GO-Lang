package main

import (
	"fmt"
	"os"
)

func main() {

	//raccessing the files
	// file, err := os.Open("example.txt")

	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fileInfo, err := file.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File name:", fileInfo.Name())
	// fmt.Println("File or Folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size()) //bytes mei retun kr rha hai
	// fmt.Println("file permission:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())

	//read file content
	file, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	//always close the file after reading it -> main function exit ho jayega tb defer function run hoga.
	defer file.Close()

	fileLen, err := file.Stat()
	//creating buffer -> buffer is array of bytes
	buff := make([]byte, fileLen.Size())

	d, err := file.Read(buff)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buff); i++ {
		fmt.Println("Data buffer", d, string(buff[i]))
	}

}
