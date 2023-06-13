package gateway

import (
	"context"

	"go.mod/domain/employer/entity"
)

type IDBProvider interface {
	GetByID(ctx context.Context, ID string) (*entity.Employer, error)
	GetByIDWithAppointments(ctx context.Context, ID string) (*entity.EmployerWithAppointmentsInfo, error)
}
