package main

import "fmt"

type Address struct {
	City  string
	State string
}
type Employee struct {
	Name    string
	Address //embedding addrress
}

func main() {
	emp := &Employee{Name: "Abhijeet", Address: Address{City: "Varanasi", State: "UP"}}
	fmt.Println(*emp)
}
