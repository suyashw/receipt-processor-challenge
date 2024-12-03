package main

import (
	"log"
	"net/http"

	"receipt-processor/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

    // Define the routes and associate them with handler functions
	r.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

    // Start the server on port 8080
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
