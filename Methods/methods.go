package main

import "fmt"

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}
type number int

type person struct {
	name string
}
type teacher struct {
	name   string
	branch string
}
type student struct {
	language string
	marks    int
}

func (n number) square() number {
	return n * n
}
func (p *person) changeName(newName string) {
	p.name = newName
}
func main() {
	details := User{"Abhijeet", "abhi@go.dev", true, 23}
	fmt.Println(details)
	details.GetStatus()
	details.NewMail()
	details.display()
	a := number(4)
	b := a.square()
	fmt.Println(b)
	detail1 := person{"Ajay"}
	fmt.Println(detail1.name)
	detail1.changeName("Abhi")
	fmt.Println(detail1.name)
	val1 := student{"Java", 50}
	val2 := teacher{"Abhi", "CSE"}
	val1.show()
	val2.show()
}
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println(u.Email)
}
func (u User) display() {
	fmt.Println("Name: ", u.Name)
	fmt.Println("Age: ", u.age)
}
func (s student) show() {
	fmt.Println("Language", s.language)
	fmt.Println("Marks", s.marks)
}
func (t teacher) show() {
	fmt.Println("Name", t.name)
	fmt.Println("Branch", t.branch)
}
