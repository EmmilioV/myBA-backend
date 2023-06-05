package entity

import "time"

type Appoinment struct {
	ID            int       `json:"id" db:"id"`
	EmployerID    string    `json:"employer_id" db:"employer_id"`
	CustomerID    string    `json:"customer_id" db:"customer_id"`
	Date          time.Time `json:"date_of" db:"date_of"`
	TotalDuration float64   `json:"total_duration" db:"total_duration"`
	TotalCost     float64   `json:"total_cost" db:"total_cost"`
}
