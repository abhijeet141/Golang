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

func updateFieldByID(obj interface{}, fieldName1 string, value1 int, fieldName2 string, value2 float64, wg *sync.WaitGroup) {
	defer wg.Done()
	value := reflect.ValueOf(obj)
	// fmt.Println(value.Kind())
	value = value.Elem() //derefrencing pointer to get slice
	// fmt.Println(value)
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		if item.FieldByName(fieldName1).Int() == int64(value1) {
			field := item.FieldByName(fieldName2)
			if field.CanSet() {
				field.SetFloat(float64(value2))
			}
		}
	}
}
func main() {
	wg := sync.WaitGroup{}
	orders := &[]Order{
		{OrderID: 205, TotalPrice: 499.99},
		{OrderID: 102, TotalPrice: 259.49},
		{OrderID: 407, TotalPrice: 749.89},
		{OrderID: 303, TotalPrice: 319.20},
		{OrderID: 108, TotalPrice: 129.99},
		{OrderID: 501, TotalPrice: 549.75},
		{OrderID: 406, TotalPrice: 999.00},
	}
	products := &[]Product{
		{ProductID: 205, TotalPrice: 499.99},
		{ProductID: 102, TotalPrice: 259.49},
		{ProductID: 407, TotalPrice: 749.89},
		{ProductID: 303, TotalPrice: 319.20},
		{ProductID: 108, TotalPrice: 129.99},
		{ProductID: 501, TotalPrice: 549.75},
		{ProductID: 406, TotalPrice: 999.00}}
	wg.Add(2)
	go updateFieldByID(orders, "OrderID", 102, "TotalPrice", 300.00, &wg)
	go updateFieldByID(products, "ProductID", 205, "TotalPrice", 500.00, &wg)
	wg.Wait()
	for _, order := range *orders {
		fmt.Printf("%+v\n", order)
	}
	for _, Product := range *products {
		fmt.Printf("%+v\n", Product)
	}
}
