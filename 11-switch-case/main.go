package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to go lang swithc case")

	dice := rand.Intn(6) + 1

	fmt.Println("Dice rolled :")

	switch dice {
	case 1:
		fmt.Println("Dice number is 1")
	case 2:
		fmt.Println("Dice number is 2")
	case 3:
		fmt.Println("Dice number is 3")
		fallthrough
	case 4:
		fmt.Println("Dice number is 4")
	case 5:
		fmt.Println("Dice number is 5")
	case 6:
		fmt.Println("Dice number is 6 and you can roll dice again")
	default:
		fmt.Println("Un-expected dice : ", dice)
	}
}
