package gateway

import (
	"context"

	"go.mod/domain/service/entity"
)

type IDBUpdater interface {
	UpdateOneByID(ctx context.Context, service *entity.Service) error
}
