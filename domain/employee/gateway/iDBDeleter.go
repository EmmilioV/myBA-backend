package gateway

import "context"

type IDBDeleter interface {
	DeleteOne(ctx context.Context, employeeID string) error
}
