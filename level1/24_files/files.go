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
	// file, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// //always close the file after reading it -> main function exit ho jayega tb defer function run hoga.
	// defer file.Close()

	// fileLen, err := file.Stat()
	// //creating buffer -> buffer is array of bytes
	// buff := make([]byte, fileLen.Size())

	// d, err := file.Read(buff)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buff); i++ {
	// 	fmt.Println("Data buffer", d, string(buff[i]))
	// }

	//another way to read the file content
	// f, err := os.ReadFile("example.txt") //readfile method only load one time the data -> not suitable if files is bigger

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(f)         // this f is data but is in buffer
	// fmt.Println(string(f)) //hello go lang

	//if we have bigger file -> streaming like in node js

	//reading folders
	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()
	// fmt.Println("current directory:", *dir) //current directory: {0xc000078180}

	// fileInfo, err := dir.ReadDir(-1) //fileInfo is slice we can loop through it

	// for _, value := range fileInfo {
	// 	fmt.Println(value.Name(), value.IsDir())
	// }

	//creating a file
	// file, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// // file.WriteString("Hey Go lang help me to build scalable application")
	// // file.WriteString(" i have appended to use web dev framework")

	// _, err = file.Write([]byte("I am here again"))
	// _, err = file.Write([]byte("I will remain at the end!"))

	// // Open the file in truncate mode to ensure only the new content is written
	// file, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // file.Write([]byte("hey go lang"))
	// // _, err = file.Seek(0, 0)
	// // file.Write([]byte("i am not here"))

	// file.WriteString("I am sahil")
	// file.Truncate(0)
	// file.WriteString("Dk")
	// file.Truncate(0)
	// file.WriteString("Darshana")

	//more method
	// file, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// bytes := []byte("hello Golang")
	// file.Write(bytes)

	//read and write to another file  (streaming file)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// //buff io package
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// //infinite loop
	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" { //EOF is the end of file
	// 			panic(err)
	// 		}
	// 		break //come out od loop
	// 	}

	// 	//writing byte
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush() //flusing the writer
	// fmt.Println("Written to new file successfully!")

	//delete a file
	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted succesfully!")
}
