package usecase

import (
	employeeGateway "go.mod/domain/employee/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type UseCases struct {
	HireEmployee   *HireEmployee
	UnhireEmployee *UnhireEmployee
}

func NewUseCases(
	employerGateways *employerGateway.Gateways,
	employeeGateways *employeeGateway.Gateways,
) *UseCases {
	return &UseCases{
		HireEmployee:   NewHireEmployee(employeeGateways, employerGateways),
		UnhireEmployee: NewUnhireEmployee(employeeGateways, employerGateways),
	}
}
