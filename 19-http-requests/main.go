package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl string = "http://localhost:3000"

func main() {

	fmt.Println("Welcome to all HTTP requests")

	// getAction := "/get-users"
	// GetHttpRequest(getAction)

	// postAction := "/get-user"
	// PostHttpRequestWithJson(postAction)

	postAction := "/get-user"
	PostHttpRequestWithFormData(postAction)
}

func GetHttpRequest(action string) {
	actionUrl := baseUrl + action

	//making the get request for give URL
	response, err := http.Get(actionUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	//decodeing the respoinse
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	stringData := string(content)

	fmt.Println("API response : ", stringData)
}

func PostHttpRequestWithJson(action string) {
	actionUrl := baseUrl + action

	requestBody := strings.NewReader(`
		{
			"id":"63edf9871b7eece855c201aa"
		}
	`)
	response, err := http.Post(actionUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var contentBody strings.Builder

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	contentResponse, _ := contentBody.Write(content)
	fmt.Println("Response Bytes ", contentResponse)

	stringResponse := contentBody.String()
	fmt.Println("String response : ", stringResponse)
}

func PostHttpRequestWithFormData(action string) {
	actionUrl := baseUrl + action

	formData := url.Values{}
	formData.Add("id", "63edf9871b7eece855c201aa")

	response, err := http.PostForm(actionUrl, formData)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var contentBody strings.Builder
	data, err := contentBody.Write(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("data byte count : ", data)
	readableResponse := contentBody.String()
	fmt.Println(readableResponse)
}
