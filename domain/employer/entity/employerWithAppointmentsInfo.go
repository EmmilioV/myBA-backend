package entity

import "go.mod/domain/appointment/entity"

type EmployerWithAppointmentsInfo struct {
	ID           string                            `json:"id"`
	Name         string                            `json:"name"`
	Appointments []*entity.AppoinmentWithExtraInfo `json:"appointments"`
}
