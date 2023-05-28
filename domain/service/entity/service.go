package entity

type Service struct {
	ID            string  `json:"id"`
	AppointmentID string  `json:"appointment_id"`
	EmployeeID    string  `json:"employee_id"`
	Type          string  `json:"type"`
	Cost          float64 `json:"cost"`
	Duration      int32   `json:"duration"`
	IsCompleted   bool    `json:"is_completed"`
	Observations  string  `json:"observations"`
}
