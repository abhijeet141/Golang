package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling")
	createFile()
}
func createFile() {
	file, err := os.Create("./file2.txt")
	checkError(err)
	writeFile(file)
	readFile(file)
	defer file.Close()
}
func writeFile(file *os.File) {
	content := "Hi there this is my file"
	len, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("Length :", len)
}
func readFile(file *os.File) {
	// data, err := os.ReadFile(fileName) //as this read the entire content of the file in the memory if file is very big then memory issue can happen
	buffer := make([]byte, 1024)
	for {
		data, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		checkError(err)
		fmt.Println(string(buffer[:data]))
	}
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
