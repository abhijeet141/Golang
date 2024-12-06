package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	maxGoRoutines := 5
	semaphore := make(chan struct{}, maxGoRoutines)
	var wg = &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			fmt.Println(i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
