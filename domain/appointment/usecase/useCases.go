package usecase

import (
	appointmentGateway "go.mod/domain/appointment/gateway"
	employerGateways "go.mod/domain/employer/gateway"
	serviceGateway "go.mod/domain/service/gateway"
)

type UseCases struct {
	AddService             *AddService
	SearchWithServicesByID *SearchWithServicesByID
}

func NewUseCases(
	serviceGateways *serviceGateway.Gateways,
	employerGateways *employerGateways.Gateways,
	appointmentGateways *appointmentGateway.Gateways,
) *UseCases {
	return &UseCases{
		AddService:             NewAddService(serviceGateways, employerGateways),
		SearchWithServicesByID: NewSearchWithServicesByID(appointmentGateways),
	}
}
