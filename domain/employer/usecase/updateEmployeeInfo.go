package usecase

import (
	"context"
	"errors"

	employeeEntity "go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type UpdateEmployeeInfo struct {
	employeeGateways *employeeGateway.Gateways
	employerGateways *employerGateway.Gateways
}

func NewUpdateEmployeeInfo(
	employeeGateways *employeeGateway.Gateways,
	employerGateways *employerGateway.Gateways,
) *UpdateEmployeeInfo {
	return &UpdateEmployeeInfo{
		employeeGateways,
		employerGateways,
	}
}

func (updateEmployeeInfo *UpdateEmployeeInfo) UseCase(
	ctx context.Context, employerID string, employee *employeeEntity.Employee,
) error {
	employer, err := updateEmployeeInfo.employerGateways.IDBProvider.GetByID(ctx, employerID)
	if err != nil {
		return err
	}

	if employer == nil {
		return errors.New("EMPLOYER_DOES_NOT_EXISTS")
	}

	return updateEmployeeInfo.employeeGateways.IDBUpdater.UpdateByID(ctx, employee)
}
