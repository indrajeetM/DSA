package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

welcome := "Welcome to go lang program";
fmt.Println(welcome)


reader := bufio.NewReader(os.Stdin);
fmt.Println("Enter rating for golang: ");

// comma ok | error ok syntax

input, _ := reader.ReadString('\n');
fmt.Println("Thank you for rating us ",input);
fmt.Printf("Type of input is %T \n",input)

}
