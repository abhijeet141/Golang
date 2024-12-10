package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func callAPI(ctx context.Context) (bool, error) {
	time.Sleep(400 * time.Millisecond)
	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("context timeout exceeded")
	}
	return true, nil
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	// userId := 12002174
	isUsersubbed, err := callAPI(ctx)
	if err != nil {
		log.Fatalf("Error fetching user status: %v", err)
	}
	if isUsersubbed {
		fmt.Println("This user is subscribed")
	}
	//calling third party api
}
