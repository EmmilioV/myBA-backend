package usecase

import (
	"context"
	"errors"

	"go.mod/domain/customer/entity"
	customerGateway "go.mod/domain/customer/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type RegisterCustomer struct {
	customerGateway  *customerGateway.Gateways
	employerGateways *employerGateway.Gateways
}

func NewRegisterCustomer(
	customerGateway *customerGateway.Gateways,
	employerGateways *employerGateway.Gateways,
) *RegisterCustomer {
	return &RegisterCustomer{
		customerGateway,
		employerGateways,
	}
}

func (registerCustomer *RegisterCustomer) UseCase(
	ctx context.Context,
	employerID string,
	customer *entity.Customer,
) error {
	employer, err := registerCustomer.employerGateways.IDBProvider.GetByID(ctx, employerID)
	if err != nil {
		return err
	}

	if employer == nil {
		return errors.New("EMPLOYER_DOES_NOT_EXISTS")
	}

	return registerCustomer.customerGateway.IDBInserter.InsertOne(ctx, customer)
}
