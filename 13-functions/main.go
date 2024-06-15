package main

import "fmt"

func main() {
	greet()
	result := addition(1, 4)
	fmt.Println("Result is ", result)

	allResult := additionInfinite(1, 3, 4, 6, 4, 3)
	fmt.Println("All result :", allResult)
}

func greet() {
	fmt.Println("Welcome to functions")
}

func addition(paramOne int, paramTwo int) int {
	return paramOne + paramTwo
}

func additionInfinite(params ...int) int {
	total := 0

	for _, value := range params {
		total += value
	}
	return total
}
