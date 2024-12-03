package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/models"
	"receipt-processor/utils"
	"sync"

	"github.com/gorilla/mux"
)

// Global data store to hold receipt IDs and their points
var (
	dataStore = make(map[string]int)
	mutex     = &sync.Mutex{}
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	// Decode the JSON request body into a Receipt object
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	points := utils.CalculatePoints(receipt)
	id := utils.GenerateUUID()

	// Store the receipt ID and its associated points in memory
	mutex.Lock()
	dataStore[id] = points
	mutex.Unlock()

	// Return the generated receipt ID as the response
	response := map[string]string{"id": id}
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Retrieve the points associated with the ID from the data store
	mutex.Lock()
	points, exists := dataStore[id]
	mutex.Unlock()

	if !exists {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Return the points as a JSON response
	response := map[string]int{"points": points}
	json.NewEncoder(w).Encode(response)
}
