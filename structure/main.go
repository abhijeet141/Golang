package main

import (
	"fmt"
	"reflect"
	"structure/mypackage"
)

type Address struct {
	Name    string
	city    string
	pincode int
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}
type Person struct {
	Name    string
	age     int
	address string
}
type Author struct {
	Name      string
	branch    string
	language  string
	particles int
}

type HR struct {
	details Author
}
type Students struct {
	personalDetail struct {
		name       string
		enrollment int
	}
	Gpa float64
}
type Person2 struct {
	Name  string
	Greet func() string
}
type Person3 struct {
	Name  string
	Greet func(string) string
}

func main() {
	fmt.Println("Structure")
	a := Address{"Abhijeet", "Varanasi", 221109}
	fmt.Println(a)
	fmt.Println(a.Name)
	emp := &Employee{"Ajay", "Yadav", 25, 100000}
	fmt.Println((*emp).firstName)
	fmt.Println(emp.lastName, emp.salary, emp.age)
	p := Person{Name: "Abhi", address: "Varanasi", age: 23}
	fmt.Println(p)
	a1 := Author{"A", "CSE", "Python", 30}
	a2 := Author{"A", "CSE", "Python", 30}
	if a1 == a2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	fmt.Println(reflect.DeepEqual(a1, a2))
	result := HR{details: Author{"A", "CSE", "Python", 30}}
	fmt.Println(result)
	st := mypackage.Student{Name: "A", Branch: "CSE", Year: 2024}
	t := mypackage.Teacher{Name: "B", Subject: "C", Exp: 10, Details: mypackage.Student{Name: "a1", Branch: "a2", Year: 10}}
	fmt.Println(st)
	fmt.Println(t)
	student := Students{personalDetail: struct {
		name       string
		enrollment int
	}{name: "Abhijeet", enrollment: 12345}}
	fmt.Println(student.personalDetail.name)
	person := Person2{
		Name: "A",
	}
	person.Greet = func() string {
		return "Hello" + person.Name
	}
	fmt.Println(person.Greet())
	person2 := Person3{
		Name: "B",
	}
	person2.Greet = func(greeting string) string {
		return greeting + "," + person2.Name
	}
}
