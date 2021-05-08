package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Groceries struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Size  string `json:"on_stock"`
}

func main() {

	http.HandleFunc("/greet", greet)

	http.HandleFunc("/groceries", getGroceries)

	http.HandleFunc("/legacy", getLegacy)

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
		{"stroopwafels", "huismerk", "2"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(groceries)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(groceries)
	}
}

func getLegacy(w http.ResponseWriter, r *http.Request) {
	groceries := []Groceries{
		{"hagelslag", "Venz", "1"},
		{"pindakaas", "Calve", "2"},
		{"stroopwafels", "huismerk", "2"},
	}
	w.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(groceries)
}
