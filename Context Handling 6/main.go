package main

import (
	"context"
	"fmt"
	"time"
)

func performTask(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Task completed sucessfully")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go performTask(ctx) //aware of the cancellation
	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	}
}
