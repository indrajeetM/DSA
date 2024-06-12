package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to rating system"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating here :")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	//following line uses the comma ok | error ok syntax
	//numRating, err := strconv.ParseFloat(input, 64)
	//The above line generates error, for not able to parse the new-line charater along with the number
	//To avoid this we will trim the input

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	//Now will check if there is any error returned
	if err != nil {
		fmt.Println("Error Occured : ", err)
	} else {
		newNum := numRating + 1
		fmt.Println("We have added 1 to your rating ", newNum)
	}

}
