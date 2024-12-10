package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Order struct {
	OrderID    int     `json:"id"`
	TotalPrice float64 `json:"price"`
}

func serialize(obj interface{}) ([]byte, error) {
	value := reflect.ValueOf(obj)
	t := value.Type()
	data := make(map[string]interface{})
	for i := 0; i < value.NumField(); i++ {
		data[t.Field(i).Name] = value.Field(i).Interface()
	}
	return json.MarshalIndent(data, "", "\t")
}
func main() {
	p := Product{ID: 101, Name: "Laptop"}
	o := Order{OrderID: 12345, TotalPrice: 499.99}
	result1, _ := serialize(p)
	result2, _ := serialize(o)
	fmt.Println(string(result1))
	fmt.Println(string(result2))
}
