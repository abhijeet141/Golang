package main

import "fmt"

type Vehicle interface {
	Drive() string
}
type car struct {
	Model string
}

type Truck struct {
	Model string
}

func (c car) Drive() string {
	return "The " + c.Model + " car"
}
func (t Truck) Drive() string {
	return "The " + t.Model + " Truck"
}
func startDriving(v Vehicle) {
	fmt.Println(v.Drive())
}
func main() {
	mycar := car{Model: "Toyota"}
	myTruck := Truck{Model: "volvo"}
	startDriving(mycar)
	startDriving(myTruck)
}
