package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// fmt.Println("Cocurent Map Read Write")
	// m := make(map[int]int)
	// wg := &sync.WaitGroup{}
	// wg.Add(3)
	// go func(wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	for i := 1; i <= 200; i++ {
	// 		r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 		m[i] = r.Intn(1000000000)
	// 	}
	// }(wg)
	// go func(wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	for key, value := range m {
	// 		fmt.Println(key, ":", value)
	// 	}
	// }(wg)
	// go func(wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	for i := 1; i <= 400; i++ {
	// 		r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 		m[i] = r.Intn(1000000000)
	// 	}
	// }(wg)
	// wg.Wait()
	fmt.Println("Cocurent Map Read Write")
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 200; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			mut.Lock()
			m[i] = r.Intn(1000000000)
			mut.Unlock()
		}
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		mut.Lock()
		for key, value := range m {
			fmt.Println(key, ":", value)
		}
		mut.Unlock()
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 400; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			mut.Lock()
			m[i] = r.Intn(1000000000)
			mut.Unlock()
		}
	}(wg)
	wg.Wait()
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
