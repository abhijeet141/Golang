package main

import (
	"fmt"
	"reflect"
)

type Details struct {
	fname   string
	lname   string
	age     int
	balance int
}

type myType string

func showDetails(id, person interface{}) {
	t1 := reflect.TypeOf(id)
	t2 := reflect.TypeOf(person)
	fmt.Println("Type1: ", t1)
	fmt.Println("Type2: ", t2)
	k1 := t1.Kind()
	k2 := t2.Kind()
	fmt.Println(k1)
	fmt.Println(k2)
	value := reflect.ValueOf(person)
	numberofFields := value.NumField()
	for i := 0; i < numberofFields; i++ {
		fmt.Println(value.Field(i))
		fmt.Println(t2.Field(i))
	}
}
func main() {
	id := myType("12002174")
	person := Details{fname: "Abhijeet", lname: "Sinha", age: 23, balance: 10000}
	showDetails(id, person)
}
