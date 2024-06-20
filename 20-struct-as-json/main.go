package main

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	FName    string `json:"first_name"`
	LName    string `json:"last_name"`
	UName    string `json:"user_name"`
	Password string `json:"-"`
	Status   bool
	Role     []string `json:"Roles,omitempty"`
}

func main() {
	fmt.Println("Welcome to struct to json")

	usersList := []Users{
		{"John", "Doe", "jd-man", "jd@123", true, []string{"admin", "user"}},
		{"Alex", "Doe", "al-man", "al@123", true, []string{"user", "contractor"}},
		{"Chris", "Doe", "cd-man", "cd@123", true, nil},
	}

	fmt.Println("Users List : ", usersList)

	//To convert the above struct into json and process it further
	//it can be done using json package and by calling Marshal methos on it.

	jsonResponse, err := json.Marshal(usersList)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Marshall json %s:", jsonResponse)

	//Below line gives more formatted resonse
	//We have to use MarshalIndent with 3 parameters
	//1. Slice list/array list
	//2. Prfix, which will be used to display the data with this given string
	//3. indent character ex, \t
	jsonResponseIndented, err := json.MarshalIndent(usersList, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	fmt.Printf("Marshall json Indented %s", jsonResponseIndented)
	fmt.Println("")

}
