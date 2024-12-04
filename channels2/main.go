package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")
	mych := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(mych chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, ischannelOpen := <-mych
		fmt.Println(val)
		fmt.Println(ischannelOpen)
		fmt.Println(<-mych)
	}(mych, wg)
	go func(mych chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(mych)
		mych <- 5
	}(mych, wg)
	wg.Wait()
}
