package usecase

import (
	appointmentGateway "go.mod/domain/appointment/gateway"
	customerGateway "go.mod/domain/customer/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type UseCases struct {
	ScheduleAnAppointment *ScheduleAnAppointment
}

func NewUseCases(
	employerGateways *employerGateway.Gateways,
	customerGateways *customerGateway.Gateways,
	appointmentGateways *appointmentGateway.Gateways,
) *UseCases {
	return &UseCases{
		ScheduleAnAppointment: NewScheduleAnAppointment(employerGateways, customerGateways, appointmentGateways),
	}
}
