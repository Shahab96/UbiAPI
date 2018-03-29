package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Test struct (Model)
type Test struct {
	Name string `json:"name"`
}

var a []Test

func getTests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(a)
}

func createTests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var test Test
	_ = json.NewDecoder(r.Body).Decode(&test)
	a = append(a, test)
	json.NewEncoder(w).Encode(a)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tests", getTests).Methods("GET")
	router.HandleFunc("/tests", createTests).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
