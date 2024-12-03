package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/models"
	"testing"

	"github.com/gorilla/mux"
)

func TestProcessReceipt(t *testing.T) {
	receipt := models.Receipt{
		Retailer:    "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:       "1.25",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		},
	}

	reqBody, _ := json.Marshal(receipt)
	req := httptest.NewRequest("POST", "/receipts/process", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ProcessReceipt(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, status)
	}

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if _, ok := response["id"]; !ok {
		t.Errorf("response should contain an ID field")
	}
}

func TestGetPoints(t *testing.T) {
	// Mock receipt ID and points
	dataStore["mock-id"] = 50

	req := httptest.NewRequest("GET", "/receipts/mock-id/points", nil) // Ensure URL matches the route definition
	rec := httptest.NewRecorder()

	// Simulate a router to match the request to the handler
	r := mux.NewRouter()
	r.HandleFunc("/receipts/{id}/points", GetPoints)
	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, status)
	}

	var response map[string]int
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if points, ok := response["points"]; !ok || points != 50 {
		t.Errorf("expected points to be 50 but got %v", points)
	}
}
