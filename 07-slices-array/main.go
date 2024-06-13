package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welome to slices")

	list := []string{}

	fmt.Println("List is : ", list)
	fmt.Printf("List type is %T:\n", list)

	//to add the items to slice you can use below
	list = append(list, "zero", "one", "two", "three", "four")
	fmt.Println("List with new values :", list)

	//you can access this list, with the given start and end index
	//FYI : the last index is never included in the list

	//start at 0 and stop 3
	fmt.Println(append(list[:3]))

	//start at 1 and stop at 4
	fmt.Println(append(list[1:4]))

	//start at 1 and stop at the last
	fmt.Println(append(list[1:]))

	//another way to create slice is using make
	newList := make([]int, 5)

	fmt.Println("new list : ", newList)

	newList[0] = 101
	newList[1] = 89
	newList[2] = 499
	newList[3] = 999
	newList[4] = 1

	fmt.Println("new list after adding values: ", newList)
	fmt.Println("Is new list sorted : ", sort.IntsAreSorted(newList))

	sort.Ints(newList)
	fmt.Println("after sorting new list : ", newList)

	fmt.Println("Is new list sorted : ", sort.IntsAreSorted(newList))

	//removing items from the list
	var index = 2

	newList = append(newList[:index], newList[index+1:]...)

	fmt.Println(" new list values after removed at 2 : ", newList)
}
