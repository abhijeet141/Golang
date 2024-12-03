package main

import "fmt"

//multiple defer statement are executed in LIFO order
func main() {
	fmt.Println("Hello")
	defer fmt.Println("Hello World") //execute at the end
	defer fmt.Println("Hello World2")
	fmt.Println("Hello2")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
