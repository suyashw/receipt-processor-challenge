package utils

import (
	"math"
	"receipt-processor/models"
	"strings"

	"github.com/google/uuid"
)

// GenerateUUID creates a new unique identifier for a receipt
func GenerateUUID() string {
	return uuid.New().String()
}

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: Alphanumeric characters in retailer name
	points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	// Rule 2 & 3: Round dollar and multiple of 0.25
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}
	if remainder := math.Mod(receipt.TotalAsFloat()*100, 25); remainder == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item description length multiple of 3
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			// Ensure correct rounding up to nearest integer
			points += int(math.Ceil(item.PriceAsFloat() * 0.2))
		}
	}

	// Rule 6: Odd day
	if receipt.IsOddDay() {
		points += 6
	}

	// Rule 7: Purchase time between 2:00pm and 4:00pm
	if receipt.IsAfternoonPurchase() {
		points += 10
	}

	return points
}
