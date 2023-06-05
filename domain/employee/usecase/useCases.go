package usecase

import (
	employeeGateway "go.mod/domain/employee/gateway"
	serviceGateway "go.mod/domain/service/gateway"
)

type UseCases struct {
	SearchByIDWithServices *SearchByIDWithServices
	UpdateServiceInfo      *UpdateServiceInfo
}

func NewUseCases(
	employeeGateways *employeeGateway.Gateways,
	serviceGateways *serviceGateway.Gateways,
) *UseCases {
	return &UseCases{
		SearchByIDWithServices: NewSearchByIDWithServices(employeeGateways),
		UpdateServiceInfo:      NewUpdateServiceInfo(employeeGateways, serviceGateways),
	}
}
