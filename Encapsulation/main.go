package main

import "fmt"

type Person struct {
	Name string
	age  int //private
}

func (p *Person) setAge(age int) {
	if age > 0 {
		p.age = age
	}
}
func (p *Person) getAge() int {
	return p.age
}
func main() {
	p := Person{Name: "Alice"}
	p.setAge(15)
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.getAge())
}
