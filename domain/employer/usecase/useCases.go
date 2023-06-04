package usecase

import (
	customerGateway "go.mod/domain/customer/gateway"
	employeeGateway "go.mod/domain/employee/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type UseCases struct {
	HireEmployee       *HireEmployee
	UnhireEmployee     *UnhireEmployee
	UpdateEmployeeInfo *UpdateEmployeeInfo
	RegisterCustomer   *RegisterCustomer
}

func NewUseCases(
	employerGateways *employerGateway.Gateways,
	employeeGateways *employeeGateway.Gateways,
	customerGateways *customerGateway.Gateways,
) *UseCases {
	return &UseCases{
		HireEmployee:       NewHireEmployee(employeeGateways, employerGateways),
		UnhireEmployee:     NewUnhireEmployee(employeeGateways, employerGateways),
		UpdateEmployeeInfo: NewUpdateEmployeeInfo(employeeGateways, employerGateways),
		RegisterCustomer:   NewRegisterCustomer(customerGateways, employerGateways),
	}
}
