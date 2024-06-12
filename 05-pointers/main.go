package main

import "fmt"

func main() {

	var ptr *int
	var ptr2 *int

	fmt.Println("Pointer 1 is : ", ptr)
	fmt.Println("Pointer 2 is : ", ptr2)

	myAge := 31

	myptr := &myAge
	fmt.Println("Pointer address :", myptr)
	fmt.Println("Pointer new value :", *myptr)

	*myptr = myAge * 7
	fmt.Println("Updateing the point value : ", *myptr)

}
