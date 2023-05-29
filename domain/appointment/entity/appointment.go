package entity

import "time"

type Appoinment struct {
	ID            int       `json:"id"`
	EmployerID    string    `json:"employer_id"`
	CustomerID    string    `json:"customer_id"`
	Date          time.Time `json:"date_of"`
	TotalDuration float64   `json:"total_duration"`
	TotalCost     float64   `json:"total_cost"`
}
