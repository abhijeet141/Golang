package main

import "fmt"

func main() {
	var p = 12
	var q = 15
	p = q
	fmt.Println(p)
	p += q
	fmt.Println(p)
	p -= q
	fmt.Println(p)
	p *= q
	fmt.Println(p)
}
