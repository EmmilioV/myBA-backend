package gateway

import (
	"context"

	"go.mod/domain/customer/entity"
)

type IDBInserter interface {
	InsertOne(ctx context.Context, employee *entity.Customer) error
}
