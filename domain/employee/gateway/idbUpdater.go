package gateway

import (
	"context"

	"go.mod/domain/employee/entity"
)

type IDBUpdater interface {
	UpdateByID(ctx context.Context, employee *entity.Employee) error
}
