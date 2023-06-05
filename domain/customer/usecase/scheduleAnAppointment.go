package usecase

import (
	"context"
	"errors"

	"go.mod/domain/appointment/entity"
	appointmentGateway "go.mod/domain/appointment/gateway"
	customerGateway "go.mod/domain/customer/gateway"
	employerGateway "go.mod/domain/employer/gateway"
)

type ScheduleAnAppointment struct {
	employerGateways    *employerGateway.Gateways
	customerGateways    *customerGateway.Gateways
	appointmentGateways *appointmentGateway.Gateways
}

func NewScheduleAnAppointment(
	employerGateways *employerGateway.Gateways,
	customerGateways *customerGateway.Gateways,
	appointmentGateways *appointmentGateway.Gateways,
) *ScheduleAnAppointment {
	return &ScheduleAnAppointment{
		employerGateways:    employerGateways,
		customerGateways:    customerGateways,
		appointmentGateways: appointmentGateways,
	}
}

func (scheduleAnAppointment *ScheduleAnAppointment) UseCase(
	ctx context.Context, appointment *entity.Appoinment,
) error {
	employer, err := scheduleAnAppointment.employerGateways.IDBProvider.GetByID(ctx, appointment.EmployerID)
	if err != nil {
		return err
	}

	if employer == nil {
		return errors.New("EMPLOYER_DOES_NOT_EXISTS")
	}

	customer, err := scheduleAnAppointment.customerGateways.IDBProvider.GetByID(ctx, appointment.CustomerID)
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.New("CUSTOMER_DOES_NOT_EXISTS")
	}

	existingAppointment, err := scheduleAnAppointment.appointmentGateways.GetOneByCustomerIDAndDate(ctx, appointment.CustomerID, appointment.Date)
	if err != nil {
		return err
	}

	if existingAppointment != nil {
		return errors.New("APPOINTMENT_ALREADY_EXISTS")
	}

	return scheduleAnAppointment.appointmentGateways.InsertOne(ctx, appointment)
}
