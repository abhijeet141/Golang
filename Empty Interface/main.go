package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func Observe(i interface{}) {
	fmt.Printf("Type is %T\n", i)
	fmt.Println("Value is", i)
}

func main() {
	value := 15.25
	name := "Abhijeet"
	Observe(value)
	Observe(name)
	var p interface{} = Person{Name: "Abhijeet", Age: 23}
	val := reflect.ValueOf(p)
	fmt.Println("Type", val.Type())
	nameField := val.FieldByName("Name")
	ageField := val.FieldByName("Age")
	fmt.Println("Name: ", nameField.String())
	fmt.Println("Age: ", ageField.Int())
}
