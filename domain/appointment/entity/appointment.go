package entity

import "time"

type Appoinment struct {
	ID            string    `json:"id"`
	EmployerID    string    `json:"employer_id"`
	CustomerID    string    `json:"customer_id"`
	Date          time.Time `json:"date"`
	TotalDuration int32     `json:"duration"`
	TotalCost     float64   `json:"total_cost"`
}
