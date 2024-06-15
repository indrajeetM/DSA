package main

import "fmt"

func main() {
	fmt.Println("Welcome to conditional")

	age := 20

	if age < 3 {
		fmt.Println("This is a baby")
	} else if age < 12 {
		fmt.Println("This is a kid")
	} else if age < 19 {
		fmt.Println("This is a teenager")
	} else if age > 60 {
		fmt.Println("This is a senior citize")
	} else {
		fmt.Println("This is a Adult")
	}
}
