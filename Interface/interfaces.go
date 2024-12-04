package main

import (
	"Interface/geometry"
	"fmt"
)

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
	r := geometry.Rect{Width: 20, Height: 10}
	c := geometry.Circle{Radius: 10}
	geometry.Measure(r)
	geometry.Measure(c)
	startDriving(mycar)
	startDriving(myTruck)
}
