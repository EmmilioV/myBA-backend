package entity

type Service struct {
	ID            int     `json:"service_id"`
	AppointmentID string  `json:"appointment_id"`
	EmployeeID    string  `json:"employee_id"`
	Type          string  `json:"type_of"`
	Cost          float64 `json:"cost_of"`
	Duration      int32   `json:"duration"`
	IsCompleted   bool    `json:"is_completed"`
	Observations  string  `json:"observations"`
}
