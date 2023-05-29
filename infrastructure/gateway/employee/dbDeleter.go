package employee

import (
	"context"

	employeeGateway "go.mod/domain/employee/gateway"
)

type DBDeleter struct{}

func NewDBDeleter() employeeGateway.IDBDeleter {
	return &DBDeleter{}
}

func (deleter *DBDeleter) DeleteOne(
	ctx context.Context, employeeID string,
) error {
	return nil
}
