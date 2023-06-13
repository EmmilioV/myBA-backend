package entity

import "time"

type AppoinmentWithExtraInfo struct {
	ID            string    `json:"appointment_id"`
	CustomerID    string    `json:"customer_id" db:"customer_id"`
	CustomerName  string    `json:"customer_name" db:"customer_name"`
	Date          time.Time `json:"date_of" db:"date_of"`
	TotalDuration float64   `json:"total_duration" db:"total_duration"`
	TotalCost     float64   `json:"total_cost" db:"total_cost"`
}
