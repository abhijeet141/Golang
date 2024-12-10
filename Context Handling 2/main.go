package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func openConnection(ctx context.Context) {
	fmt.Println("Attempting Connection...")
	if rand.Intn(100) > 50 {
		time.Sleep(4 * time.Millisecond)
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Context Deadline Exceeded")
		} else {
			fmt.Println("Hanging Connection")
		}

	} else {
		time.Sleep(2 * time.Millisecond)
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Context Deadline Exceeded")
		} else {
			fmt.Println("Connection Established")
		}

	}

}
func openConnectionWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	openConnection(ctx)

}
func main() {
	openConnectionWithTimeout()
}
