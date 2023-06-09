package usecase

import (
	"context"
	"errors"

	"go.mod/domain/common"
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

	err = updateServiceInfo.serviceGateways.IDBUpdater.UpdateOneByID(ctx, service)
	if err != nil {
		return err
	}

	event := &common.Event{
		Action: "ServiceUpdated",
		UserId: employeeID,
		New:    service,
	}

	return updateServiceInfo.serviceGateways.IMQPublisher.ServiceUpdated(ctx, event)
}
