package gateway

import (
	"context"

	"go.mod/domain/employee/entity"
)

type IDBInserter interface {
	InsertOne(ctx context.Context, employee *entity.Employee) error
}
