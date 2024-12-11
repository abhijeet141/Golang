package main

import (
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func filterData(obj interface{}, minAge int, maxAge int, fieldName string) []map[string]interface{} {
	value := reflect.ValueOf(obj)
	var res []map[string]interface{}
	for i := 0; i < value.Len(); i++ {
		index := value.Index(i)
		fieldValue := index.FieldByName(fieldName)
		if !fieldValue.IsValid() {
			panic("Invalid field name")
		}
		if fieldValue.Int() >= int64(minAge) && fieldValue.Int() <= int64(maxAge) {
			structMap := make(map[string]interface{})
			for j := 0; j < index.NumField(); j++ {
				fieldName := index.Type().Field(j).Name
				structMap[fieldName] = index.Field(j).Interface()
				log.Println(fieldName)
				log.Println(structMap)
			}
			res = append(res, structMap)
		} else {
			panic("Error")
		}
		// 	structMap := make(map[string]interface{})
		// 	for j := 0; j < index.NumField(); j++ {
		// 		fieldName := index.Type().Field(j).Name
		// 		structMap[fieldName] = index.Field(j).Interface()
		// 	}
		// 	res = append(res, structMap)
		// } else {
		// 	log.Println(res)
		// 	panic("Field Type error")
		// }
	}
	return res
}
func main() {
	users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
		{Name: "Diana", Age: 28},
	}
	data := filterData(users, 15, 25, "Age")
	fmt.Println(data)
}
