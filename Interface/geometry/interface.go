package geometry

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perimerter() float64
}

type Rect struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (R Rect) Area() float64 {
	return R.Width * R.Height
}
func (C Circle) Area() float64 {
	return C.Radius * C.Radius * math.Pi
}
func (R Rect) Perimerter() float64 {
	return 2*R.Width + 2*R.Height
}
func (C Circle) Perimerter() float64 {
	return 2 * math.Pi * C.Radius
}
func Measure(G Geometry) {
	fmt.Println(G)
	fmt.Println(G.Area())
	fmt.Println(G.Perimerter())
}
