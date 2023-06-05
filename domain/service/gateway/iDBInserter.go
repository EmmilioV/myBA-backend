package gateway

import (
	"context"

	"go.mod/domain/service/entity"
)

type IDBInserter interface {
	InsertOne(ctx context.Context, service *entity.Service) error
}
