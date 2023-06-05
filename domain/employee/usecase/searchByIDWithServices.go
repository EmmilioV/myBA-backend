package usecase

import (
	"context"

	"go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
)

type SearchByIDWithServices struct {
	employeeGateways *employeeGateway.Gateways
}

func NewSearchByIDWithServices(
	employeeGateways *employeeGateway.Gateways,
) *SearchByIDWithServices {
	return &SearchByIDWithServices{
		employeeGateways: employeeGateways,
	}
}

func (searchByIDWithServices *SearchByIDWithServices) UseCase(
	ctx context.Context, employeeID string,
) (*entity.EmployeeWithServicesInfo, error) {
	return searchByIDWithServices.employeeGateways.IDBProvider.GetByIDWithServicesInformation(ctx, employeeID)
}
