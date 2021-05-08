package app

import (
	"log"
	"net/http"
)

func Start() {
	log.Println("Starting App service")
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/groceries", getGroceries)
	http.HandleFunc("/legacy", getLegacy)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
