package main

import "fmt"

func main() {

	fmt.Println("Welcome to array")

	var numList [4]string

	numList[0] = "one"
	numList[1] = "two"
	numList[2] = "three"
	numList[3] = "four"

	fmt.Println(numList)

	colorList := [3]string{"red", "green", "blue"}

	fmt.Println(colorList)

	var friends [3]string

	friends[0] = "one"
	friends[2] = "three"

	fmt.Println(friends)

	friends[1] = "two"
	fmt.Println(friends)

}
