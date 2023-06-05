package gateway

import (
	"context"

	"go.mod/domain/service/entity"
)

type IDBProvider interface {
	GetByID(ctx context.Context, ID string) (*entity.Service, error)
	GetAll(ctx context.Context) ([]*entity.Service, error)
}
