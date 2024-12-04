package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func someFunc(num string) {
	defer wg.Done()
	fmt.Println(num)
}
func main() {
	wg.Add(3)
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")
	// time.Sleep(3 * time.Millisecond)
	wg.Wait()
	fmt.Println("Hi")
}
