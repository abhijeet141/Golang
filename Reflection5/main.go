package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

type Order struct {
	OrderID    int
	TotalPrice float64
}

type Product struct {
	ProductID  int
	TotalPrice float64
}

func sortbyId(obj interface{}, fieldName string, wg *sync.WaitGroup) {
	defer wg.Done()
	value := reflect.ValueOf(obj)
	sort.Slice(obj, func(i, j int) bool {
		return value.Index(i).FieldByName(fieldName).Int() < value.Index(j).FieldByName(fieldName).Int()
	})
}
func main() {
	wg := &sync.WaitGroup{}
	orders := []Order{
		{OrderID: 205, TotalPrice: 499.99},
		{OrderID: 102, TotalPrice: 259.49},
		{OrderID: 407, TotalPrice: 749.89},
		{OrderID: 303, TotalPrice: 319.20},
		{OrderID: 108, TotalPrice: 129.99},
		{OrderID: 501, TotalPrice: 549.75},
		{OrderID: 406, TotalPrice: 999.00},
	}
	products := []Product{
		{ProductID: 200, TotalPrice: 499.99},
		{ProductID: 100, TotalPrice: 259.49},
		{ProductID: 400, TotalPrice: 749.89},
		{ProductID: 300, TotalPrice: 319.20},
		{ProductID: 100, TotalPrice: 129.99},
		{ProductID: 500, TotalPrice: 549.75},
		{ProductID: 400, TotalPrice: 999.00}}
	wg.Add(2)
	go sortbyId(orders, "OrderID", wg)
	go sortbyId(products, "ProductID", wg)
	wg.Wait()
	for _, value := range orders {
		fmt.Println(value)
	}
	fmt.Println("----")
	for _, value := range products {
		fmt.Println(value)
	}
}
