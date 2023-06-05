package gateway

import (
	"context"

	"go.mod/domain/appointment/entity"
)

type IDBInserter interface {
	InsertOne(ctx context.Context, appointment *entity.Appoinment) error
}
