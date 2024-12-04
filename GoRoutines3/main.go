package main

import (
	"fmt"
	"sync"
)

func myfun(value string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(value, ":", 0)
	}
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go myfun("goRoutine", wg)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("Going")
	wg.Wait()
	fmt.Println("done")
}
