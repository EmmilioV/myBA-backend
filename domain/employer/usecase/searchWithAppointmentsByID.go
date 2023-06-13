package usecase

import (
	"context"
	"errors"

	employerEntity "go.mod/domain/employer/entity"
	employerGateway "go.mod/domain/employer/gateway"
)

type SearchWithAppointmentsByID struct {
	employerGateways *employerGateway.Gateways
}

func NewsearchWithAppointmentsByID(
	employerDBProvider *employerGateway.Gateways,
) *SearchWithAppointmentsByID {
	return &SearchWithAppointmentsByID{
		employerGateways: employerDBProvider,
	}
}

func (searchAppointmentByID *SearchWithAppointmentsByID) UseCase(
	ctx context.Context, employeerID string,
) (*employerEntity.EmployerWithAppointmentsInfo, error) {
	employer, err := searchAppointmentByID.employerGateways.IDBProvider.GetByID(ctx, employeerID)
	if err != nil {
		return nil, err
	}

	if employer == nil {
		return nil, errors.New("EMPLOYER_NOT_FOUND")
	}

	return searchAppointmentByID.employerGateways.IDBProvider.GetByIDWithAppointments(ctx, employeerID)
}
