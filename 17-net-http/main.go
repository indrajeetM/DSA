package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://indrajeet.tech"

func main() {
	fmt.Println("Welcome to http requests")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occured")
		panic(err)
	}

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occured in decoding response")
		panic(err)
	}
	data := string(databytes)
	fmt.Println("Data received : ", data)
}
