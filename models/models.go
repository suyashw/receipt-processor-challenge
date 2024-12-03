package models

import (
	"strconv"
	"time"
)

// Receipt represents the structure of a receipt
type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Total        string  `json:"total"`
	Items        []Item  `json:"items"`
}

// Item represents an item in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// TotalAsFloat converts the total amount to a float64
func (r *Receipt) TotalAsFloat() float64 {
	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return 0
	}
	return total
}

// PriceAsFloat converts the price of an item to a float64
func (i *Item) PriceAsFloat() float64 {
	price, err := strconv.ParseFloat(i.Price, 64)
	if err != nil {
		return 0
	}
	return price
}

// IsOddDay checks if the purchase day is odd
func (r *Receipt) IsOddDay() bool {
	date, err := time.Parse("2006-01-02", r.PurchaseDate)
	if err != nil {
		return false
	}
	return date.Day()%2 != 0
}

// IsAfternoonPurchase checks if the purchase time is between 2:00pm and 4:00pm
func (r *Receipt) IsAfternoonPurchase() bool {
	timeOfDay, err := time.Parse("15:04", r.PurchaseTime)
	if err != nil {
		return false
	}
	return timeOfDay.Hour() == 14 || (timeOfDay.Hour() == 15 && timeOfDay.Minute() < 60)
}
