package usecase

import (
	"context"
	"errors"

	employerGateways "go.mod/domain/employer/gateway"
	"go.mod/domain/service/entity"
	serviceGateway "go.mod/domain/service/gateway"
)

type AddService struct {
	serviceGateways  *serviceGateway.Gateways
	employerGateways *employerGateways.Gateways
}

func NewAddService(
	serviceGateways *serviceGateway.Gateways,
	employerGateways *employerGateways.Gateways,
) *AddService {
	return &AddService{
		serviceGateways:  serviceGateways,
		employerGateways: employerGateways,
	}
}

func (addService *AddService) UseCase(
	ctx context.Context, service *entity.Service, employerID string,
) error {
	employer, err := addService.employerGateways.IDBProvider.GetByID(ctx, employerID)
	if err != nil {
		return err
	}

	if employer == nil {
		return errors.New("EMPLOYER_DOES_NOT_EXISTS")
	}

	return addService.serviceGateways.IDBInserter.InsertOne(ctx, service)
}
