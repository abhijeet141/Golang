package main

import (
	"context"
	"fmt"
)

func getUser(ctx context.Context) {
	if userID := ctx.Value("userID"); userID != nil {
		fmt.Println("UserID: ", userID)
	} else {
		fmt.Println("User Id not found")
	}
}
func main() {
	ctx := context.WithValue(context.Background(), "userID", 42)
	getUser(ctx)
}
