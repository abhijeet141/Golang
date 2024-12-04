package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer
var signals = []string{"test"}

func main() {
	go greeter("Hello") //we never waited our thread to comeback
	greeter("World")
	websiteList := []string{
		"https://google.com",
		"https://github.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait() //responsible for not allowing main method to finish
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s)
	}

}

func getStatusCode(endPoint string) {
	defer wg.Done()
	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Println("something wrong")
	} else {
		mut.Lock()
		signals = append(signals, endPoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endPoint)
	}
}
