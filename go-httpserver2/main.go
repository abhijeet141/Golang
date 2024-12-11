package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func multiplyNumbers(num1 int, num2 int) int {
	return num1 * num2
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Println("App started and listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server", err)
	}

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	numStr1 := r.URL.Query().Get("num1")
	numStr2 := r.URL.Query().Get("num2")
	num1, err1 := strconv.Atoi(numStr1)
	num2, err2 := strconv.Atoi(numStr2)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid Number", http.StatusBadRequest)
	}
	ans := multiplyNumbers(num1, num2)
	response := fmt.Sprintf("The Product is %d", ans)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(response))
}
