package gateway

import (
	"context"

	"go.mod/domain/employee/entity"
)

type IDBProvider interface {
	GetByIDWithServicesInformation(ctx context.Context, employeeID string) (*entity.EmployeeWithServicesInfo, error)
}
