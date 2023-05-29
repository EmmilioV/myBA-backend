package usecase

import (
	"context"
	"errors"

	"go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type HireEmployee struct {
	employeeGateways *employeeGateway.Gateways
	employerGateways *employerGateway.Gateways
}

func NewHireEmployee(
	employeeGateways *employeeGateway.Gateways,
	employerGateways *employerGateway.Gateways,
) *HireEmployee {
	return &HireEmployee{
		employeeGateways,
		employerGateways,
	}
}

func (hireEmployee *HireEmployee) UseCase(
	ctx context.Context,
	employerID string,
	employee *entity.Employee,
) error {
	employer, err := hireEmployee.employerGateways.IDBProvider.GetByID(ctx, employerID)
	if err != nil {
		return err
	}

	if employer == nil {
		return errors.New("EMPLOYER_DOES_NOT_EXISTS")
	}

	return hireEmployee.employeeGateways.IDBInserter.InsertOne(ctx, employee)
}
