package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	EncodeJson()
}
func EncodeJson() {
	courses := []course{
		{"ReactJs", 299, "code.in", "abc123", []string{"web", "js"}},
		{"ExpressJs", 299, "code.in", "abc123", []string{"web", "js"}},
		{"MERN", 299, "code.in", "abc123", nil}}
	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
	DecodeJson(string(finalJson))
}
func DecodeJson(jsonData string) {
	fmt.Println(jsonData)
	var courses []course
	checkValid := json.Valid([]byte(jsonData))
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal([]byte(jsonData), &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		panic("Json was not valid ")
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonData), &data)
	fmt.Printf("%#v\n", data)
	for key, value := range data {
		fmt.Printf("Key is %v and value is %v\n", key, value)
	}
}
