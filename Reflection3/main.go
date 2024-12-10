package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	jsonData := `{"name":"Abhi","age":23}`
	var p Person
	err := json.Unmarshal([]byte(jsonData), &p)
	checkError(err)
	v := reflect.ValueOf(p)
	fmt.Println(p)
	fmt.Println(v)
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i))
	}
	var data map[string]interface{}
	err1 := json.Unmarshal([]byte(jsonData), &data)
	checkError(err1)
	fmt.Printf("%#v\n", data)
	for key, value := range data {
		fmt.Printf("Key is %v and value is %v\n", key, value)
	}
}
