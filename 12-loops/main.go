package main

import "fmt"

func main() {
	fmt.Println("welcome to loop")

	list := []string{"one", "two", "three"}

	fmt.Println("Item in list ", list)

	//THe below line code works but, not as optimized and not commany used
	// for i := 0; i < len(list); i++ {
	// 	fmt.Printf(" value at index %v : %v\n", i, list[i])
	// }

	for index, value := range list {
		fmt.Printf("Index is %v and value is %v\n", index, value)
	}
}
