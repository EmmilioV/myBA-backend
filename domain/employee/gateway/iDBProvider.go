package gateway

import (
	"context"

	"go.mod/domain/employee/entity"
)

type IDBProvider interface {
	GetByID(ctx context.Context, employeeID string) (*entity.Employee, error)
	GetByIDWithServicesInformation(ctx context.Context, employeeID string) (*entity.EmployeeWithServicesInfo, error)
}
