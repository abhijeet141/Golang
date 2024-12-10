package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func callInventoryService(ctx context.Context) string {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://inventory.example.com", nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	fmt.Println(res.StatusCode)
	requestId := ctx.Value("requestID").(string)
	fmt.Println(requestId)
	return "Product Available"
}
func callShippingService(ctx context.Context) string {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://shipping.example.com", nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	fmt.Println(res.StatusCode)
	userID := ctx.Value("userID").(string)
	fmt.Println(userID)
	return "Shipping Details Calculated"
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "requestID", "12345")
	ctx = context.WithValue(ctx, "userID", "user_789")
	inventoryResult := callInventoryService(ctx)
	shippingResult := callShippingService(ctx)
	fmt.Println("Inventory Result:", inventoryResult)
	fmt.Println("Shipping Result:", shippingResult)
}
