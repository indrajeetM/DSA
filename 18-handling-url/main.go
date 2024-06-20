package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://localhost:3000/get-user-data?ip=127.0.0.1&location=remote"

func main() {

	fmt.Println("Welcome to url handling")
	fmt.Println("URL :", myUrl)

	result, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error in parsing url")
		panic(err)
	}

	fmt.Println("This is parsed URL : ", result)
	fmt.Println("This is URL param: ", result.Host)
	fmt.Println("This is URL port: ", result.Port())
	fmt.Println("This is URL path: ", result.Path)
	fmt.Println("This is URL scheme: ", result.Scheme)
	fmt.Println("This is URL query: ", result.Query())

	queryParams := result.Query()

	for index, value := range queryParams {
		fmt.Println("Index : ", index)
		fmt.Println("Value : ", value)
	}

}
