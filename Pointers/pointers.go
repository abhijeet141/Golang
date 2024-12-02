package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)
	myNumber := 10
	var myptr *int = &myNumber
	fmt.Println(myptr)
	fmt.Println(*myptr)
}
