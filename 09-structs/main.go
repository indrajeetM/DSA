package main

import "fmt"

func main() {
	fmt.Println("This is structur aka clases")

	johnDoe := User{"John Doe", 22, "jd@gmail.com", "Software engineer"}

	//below line prints the values without the keys
	fmt.Println(johnDoe)

	//below line prints the value with the keys
	fmt.Printf("The details are %+v\n", johnDoe)
}

type User struct {
	Name  string
	Age   int
	Email string
	job   string
}
