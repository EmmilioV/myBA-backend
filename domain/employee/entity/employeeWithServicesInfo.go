package entity

import "go.mod/domain/service/entity"

type EmployeeWithServicesInfo struct {
	ID       string                              `json:"id" db:"employee_id"`
	Name     string                              `json:"name" db:"employee_name"`
	Email    string                              `json:"email" db:"employee_email"`
	Services []*entity.ServiceWithAdditionalInfo `json:"services"`
}
