package usecase

import (
	"context"
	"errors"

	employeeGateway "go.mod/domain/employee/gateway"
	"go.mod/domain/service/entity"
	serviceGateway "go.mod/domain/service/gateway"
)

type UpdateServiceInfo struct {
	employeeGateways *employeeGateway.Gateways
	serviceGateways  *serviceGateway.Gateways
}

func NewUpdateServiceInfo(
	employeeGateways *employeeGateway.Gateways,
	serviceGateways *serviceGateway.Gateways,
) *UpdateServiceInfo {
	return &UpdateServiceInfo{
		employeeGateways,
		serviceGateways,
	}
}

func (updateServiceInfo *UpdateServiceInfo) UseCase(
	ctx context.Context, employeeID string, service *entity.Service,
) error {
	employee, err := updateServiceInfo.employeeGateways.IDBProvider.GetByID(ctx, employeeID)
	if err != nil {
		return err
	}

	if employee == nil {
		return errors.New("EMPLOYEE_DOES_NOT_EXISTS")
	}

	return updateServiceInfo.serviceGateways.IDBUpdater.UpdateOneByID(ctx, service)
}
