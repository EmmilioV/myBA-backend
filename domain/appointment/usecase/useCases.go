package usecase

import (
	employerGateways "go.mod/domain/employer/gateway"
	serviceGateway "go.mod/domain/service/gateway"
)

type UseCases struct {
	AddService *AddService
}

func NewUseCases(
	serviceGateways *serviceGateway.Gateways,
	employerGateways *employerGateways.Gateways,
) *UseCases {
	return &UseCases{
		AddService: NewAddService(serviceGateways, employerGateways),
	}
}
