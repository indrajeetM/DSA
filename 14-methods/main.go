package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in classes aka Struct")

	jd := Users{Name: "john Doe", Age: 33, MaritialStatus: false, Email: "jd@gmail.com"}
	fmt.Println("User JD : ", jd)

	jd.GetStaus()
	fmt.Println("User Before Email : ", jd.Email)
	jd.NewEmail()
	fmt.Println("User After Email : ", jd.Email)

}

type Users struct {
	Name           string
	Age            int
	MaritialStatus bool
	Email          string
}

func (u Users) GetStaus() {
	fmt.Println("User maritial status : ", u.MaritialStatus)
}

func (u Users) NewEmail() {
	u.Email = "new-jd@gmail.com"
	fmt.Println("User email is : ", u.Email)
}
