package main

import (
	"fmt"
	"reflect"
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

func findbyID(obj interface{}, fieldName string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	value := reflect.ValueOf(obj)
	for i := 0; i < value.Len(); i++ {
		if value.Index(i).FieldByName(fieldName).Int() == int64(id) {
			fmt.Println(value.Index(i))
		} else {
			continue
		}
	}
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
		{ProductID: 205, TotalPrice: 499.99},
		{ProductID: 102, TotalPrice: 259.49},
		{ProductID: 407, TotalPrice: 749.89},
		{ProductID: 303, TotalPrice: 319.20},
		{ProductID: 108, TotalPrice: 129.99},
		{ProductID: 501, TotalPrice: 549.75},
		{ProductID: 406, TotalPrice: 999.00}}
	wg.Add(2)
	go findbyID(orders, "OrderID", 303, wg)
	go findbyID(products, "ProductID", 406, wg)
	wg.Wait()
}
