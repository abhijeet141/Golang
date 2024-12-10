package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	employees := []Employee{
		{ID: 1, Name: "Abhijeet", Age: 30, Salary: 50000},
		{ID: 1, Name: "Ajay", Age: 25, Salary: 500000},
		{ID: 3, Name: "Aditya", Age: 25, Salary: 10000},
	}
	jsonData, err := json.MarshalIndent(employees, "", "\t")
	checkError(err)
	fmt.Println("JSON DATA")
	fmt.Println(string(jsonData))
	jsonInput := `[
        {
                "id": 1,
                "name": "Abhijeet",
                "age": 30,
                "salary": 50000
        },
        {
                "id": 1,
                "name": "Ajay",
                "age": 25,
                "salary": 500000
        },
        {
                "id": 3,
                "name": "Aditya",
                "age": 25,
                "salary": 10000
        }
			]`
	var newEmployee []Employee
	err = json.Unmarshal([]byte(jsonInput), &newEmployee)
	checkError(err)
	for _, emp := range newEmployee {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Salary: %d\n", emp.ID, emp.Name, emp.Age, emp.Salary)
	}
}
