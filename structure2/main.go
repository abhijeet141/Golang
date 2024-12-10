package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	c := map[string]*Employee{} //empty map
	e := Employee{
		Name: "Abhi", Number: 1, Hired: time.Now(),
	}
	b := Employee{"Ajay", 2, nil, time.Now()}
	e.Boss = &b
	fmt.Println(e)
	fmt.Println(b)
	c["Atul"] = &Employee{"Atul", 2, nil, time.Now()}
	c["Aman"] = &Employee{Name: "Aman", Number: 1, Boss: c["Atul"], Hired: time.Now()}
}
