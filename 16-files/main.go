package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to file systems, writing and reading from file")
	fileContent := "This is a file content"
	filePath := "myFile.txt"

	//below line just creates the file if not existed only in write mode.
	// file, err := os.Create(filePath)

	//below line creates a file if does not exitst and opens it in append mode
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error occured while creating file :")
		panic(err)
	}
	defer file.Close()

	length, err := io.WriteString(file, fileContent)
	if err != nil {
		fmt.Println("Error occured while writing into file")
		panic(err)
	}

	fmt.Println("Lenght of file after writing content:", length)

	readFileContent(filePath)

}

func readFileContent(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error occured while reading file content")
		panic(err)
	}
	fmt.Println("Printing raw file content : ", content)

	convertedToString := string(content)
	fmt.Println("Converted file content : ", convertedToString)
}
