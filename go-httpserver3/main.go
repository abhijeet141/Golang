package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
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
		if int(fieldValue.Int()) >= minAge && int(fieldValue.Int()) <= maxAge {
			structMap := make(map[string]interface{})
			for j := 0; j < index.NumField(); j++ {
				fieldName := index.Type().Field(j).Name
				structMap[fieldName] = index.Field(j).Interface()
			}
			res = append(res, structMap)
		}
	}
	return res
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	log.Println("App started and running on PORT 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
		{Name: "Diana", Age: 28},
	}
	minAgeStr := r.URL.Query().Get("minAge")
	maxAgeStr := r.URL.Query().Get("maxAge")
	minAge, err1 := strconv.Atoi(minAgeStr)
	maxAge, err2 := strconv.Atoi(maxAgeStr)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid Number", http.StatusBadRequest)
		return
	}
	data := filterData(users, minAge, maxAge, "Age")
	dataJson, err := json.MarshalIndent(map[string]interface{}{
		"Users": data}, "", "\t")
	if err != nil {
		http.Error(w, "Error Marshalling Data", http.StatusInternalServerError)
		return
	}
	// response := fmt.Sprintln(dataJson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataJson)
}
