package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Requests")
	PerformGetRequest()
	PerformPostRequest()
}
func PerformGetRequest() {
	const myurl = "https://localhost:800/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	var responseString strings.Builder
	data, err := io.ReadAll(response.Body)
	bytecount, _ := responseString.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytecount)
	fmt.Println(responseString.String())
	// fmt.Println(string(data))
}
func PerformPostRequest() {
	const myurl = "https://localhost:800/post"
	requestBody := strings.NewReader(`
	{
		"coursename":"Golang",
		"price":0
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
