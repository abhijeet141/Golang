package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web Requests")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response %T\n", res)
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println(content)
}
