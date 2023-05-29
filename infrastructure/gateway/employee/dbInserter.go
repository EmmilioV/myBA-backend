package employee

import (
	"context"

	"go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
)

type DBInserter struct{}

func NewDBInserter() employeeGateway.IDBInserter {
	return &DBInserter{}
}

func (inserter *DBInserter) InsertOne(
	ctx context.Context, employee *entity.Employee,
) error {
	return nil
}
