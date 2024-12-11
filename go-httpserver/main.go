package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func calculateSum(num int) int {
	ans := 0
	for i := 1; i <= num; i++ {
		ans += i
	}
	return ans
}
func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/{id}", HomeHandler).Methods("GET") //path parameter
	r.HandleFunc("/", HomeHandler).Methods("GET")
	log.Println("App started and listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}

// passing a pointer avoid copying the entire request struct, modify the request if needed
// responsewriter is an interface and they are already reference types
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//path parameter
	// vars := mux.Vars(r)
	// idStr := vars["id"]
	//query Parameter
	idStr := r.URL.Query().Get("id")
	// fmt.Println(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	ans := calculateSum(id)
	response := fmt.Sprintf("Hi your answer for the algorithm is %d", ans)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(response))
}
