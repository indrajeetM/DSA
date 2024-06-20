package main

import "fmt"

func main() {

	fmt.Println("This is the first print statement")

	// Defere executes the line code at the very last of the  function execution.
	// if there is only defer in the program it will be executed right before the end of program
	// if there are multiple defer statements, then they will be executed in LIFO manner.
	// i.e statement which is mentioned very last will be first one to execute.

	defer fmt.Println("101")
	defer fmt.Println("102")
	defer fmt.Println("103")

	//If you comment the Line: 25 i.e. function call LoopOver()
	//The above line of code prints 103,102,101

	//Now un-comment the Line: 25,.
	//In the function description it has defer used to print the index value.
	//Check the flow now, while printing the index value in loop with defer statement.
	//And the output would be 4321 103 102 101
	// LoopOver()
	fmt.Println("")

}

func LoopOver() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
