package gateway

import (
	"context"

	"go.mod/domain/customer/entity"
)

type IDBProvider interface {
	GetByID(ctx context.Context, ID string) (*entity.Customer, error)
}
