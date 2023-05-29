package usecase

import (
	"context"
	"errors"

	employeeGateway "go.mod/domain/employee/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type UnhireEmployee struct {
	employeeGateways *employeeGateway.Gateways
	employerGateways *employerGateway.Gateways
}

func NewUnhireEmployee(
	employeeGateways *employeeGateway.Gateways,
	employerGateways *employerGateway.Gateways,
) *UnhireEmployee {
	return &UnhireEmployee{
		employeeGateways,
		employerGateways,
	}
}

func (unhireEmployee *UnhireEmployee) UseCase(
	ctx context.Context,
	employerID string,
	employeeID string,
) error {
	employer, err := unhireEmployee.employerGateways.IDBProvider.GetByID(ctx, employerID)
	if err != nil {
		return err
	}

	if employer == nil {
		return errors.New("EMPLOYER_DOES_NOT_EXISTS")
	}

	return unhireEmployee.employeeGateways.IDBDeleter.DeleteOne(ctx, employeeID)
}
