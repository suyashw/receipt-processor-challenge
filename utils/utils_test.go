package utils

import (
	"receipt-processor/models"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	id := GenerateUUID()
	if len(id) == 0 {
		t.Errorf("expected a valid UUID but got an empty string")
	}
}

func TestCalculatePoints(t *testing.T) {
	receipt := models.Receipt{
		Retailer:    "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "14:33",
		Total:       "9.00",
		Items: []models.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 107
	if points != expectedPoints {
		t.Errorf("expected points to be %d but got %d", expectedPoints, points)
	}
}
