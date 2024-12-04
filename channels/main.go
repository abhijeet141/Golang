package main

import "fmt"

func myFunc(mych chan int) {
	fmt.Println(200 + <-mych)
}
func main() {
	fmt.Println("Channels")
	mych := make(chan int)
	go myFunc(mych)
	mych <- 5
	fmt.Println("Hello")
}
