package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Groceries struct {
	Name  string
	Brand string
	Size  string
}

func main() {

	http.HandleFunc("/greet", greet)

	http.HandleFunc("/groceries", getGroceries)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
	log.Println("Request accepted")
}

func getGroceries(w http.ResponseWriter, r *http.Request) {
	groceries := []Groceries{
		{"hagelslag", "Venz", "1"},
		{"pindakaas", "Calve", "2"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groceries)
}
