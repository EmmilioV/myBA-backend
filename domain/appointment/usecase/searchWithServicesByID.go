package usecase

import (
	"context"

	"go.mod/domain/appointment/entity"
	appointmentGateway "go.mod/domain/appointment/gateway"
)

type SearchWithServicesByID struct {
	appointmentGateways *appointmentGateway.Gateways
}

func NewSearchWithServicesByID(
	appointmentGateways *appointmentGateway.Gateways,
) *SearchWithServicesByID {
	return &SearchWithServicesByID{
		appointmentGateways: appointmentGateways,
	}
}

func (searchAllServicesByID *SearchWithServicesByID) UseCase(
	ctx context.Context, appointmentID string,
) (*entity.AppointmentWithServicesInfo, error) {
	return searchAllServicesByID.appointmentGateways.IDBProvider.GetWithServicesByID(ctx, appointmentID)
}
