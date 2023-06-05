package entity

import "time"

type ServiceWithAdditionalInfo struct {
	ID           int       `json:"service_id"`
	CustomerName string    `json:"customer_name"`
	Date         time.Time `json:"date_of"`
	Type         string    `json:"type_of"`
	Cost         float64   `json:"cost_of"`
	Duration     int32     `json:"duration"`
	IsCompleted  bool      `json:"is_completed"`
	Observations string    `json:"observations"`
}
