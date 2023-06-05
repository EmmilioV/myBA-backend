package usecase

import employeeGateway "go.mod/domain/employee/gateway"

type UseCases struct {
	SearchByIDWithServices *SearchByIDWithServices
}

func NewUseCases(
	employeeGateways *employeeGateway.Gateways,
) *UseCases {
	return &UseCases{
		SearchByIDWithServices: NewSearchByIDWithServices(employeeGateways),
	}
}
