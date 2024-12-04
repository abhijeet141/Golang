package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mych := make(chan string)
	ch := make(chan string)
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		mych <- "data"
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- "otherdata"
	}(wg)
	go func() {
		wg.Wait()
		close(mych)
		close(ch)
	}()
	select {
	case message1 := <-mych:
		fmt.Println(message1)
	case message2 := <-ch:
		fmt.Println(message2)
	}
}
