package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	OrderID    int
	TotalPrice float64
}
type Product struct {
	ProductID  int
	TotalPrice float64
}

func fieldStatistics(obj interface{}, fieldName string) {
	value := reflect.ValueOf(obj)
	value = value.Elem()
	func(obj interface{}, fieldName string) {
		var max float64
		for i := 0; i < value.Len(); i++ {
			index := value.Index(i)
			if index.FieldByName(fieldName).Float() > max {
				max = index.FieldByName(fieldName).Float()
			}
		}
		fmt.Println("Maximum Price is: ", max)
	}(obj, fieldName)
	func(obj interface{}, fieldName string) {

	}(obj, fieldName)
}
func main() {
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
		{ProductID: 406, TotalPrice: 1324.20}}
	fieldStatistics(orders, "TotalPrice")
	fieldStatistics(products, "TotalPrice")

}
