package gateway

import (
	"context"
	"time"

	"go.mod/domain/appointment/entity"
)

type IDBProvider interface {
	GetOneByCustomerIDAndDate(ctx context.Context, customerID string, date time.Time) (*entity.Appoinment, error)
	GetWithServicesByID(ctx context.Context, ID string) (*entity.AppointmentWithServicesInfo, error)
}
