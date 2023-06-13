package appointment

import (
	"context"
	"fmt"
	"log"
	"time"

	appointmentEntity "go.mod/domain/appointment/entity"
	appointmentGateway "go.mod/domain/appointment/gateway"
	serviceEntity "go.mod/domain/service/entity"
	db "go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *db.Connection
}

func NewDBProvider(
	dbConnection *db.Connection,
) appointmentGateway.IDBProvider {
	return &DBProvider{
		DBConnection: dbConnection,
	}
}

func (provider *DBProvider) GetOneByCustomerIDAndDate(
	ctx context.Context, customerID string, date time.Time,
) (*appointmentEntity.Appoinment, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE customer_id=$1 AND date_of=$2", db.Appointment)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, customerID, date)
	if err != nil {
		return nil, err
	}

	var appointment *appointmentEntity.Appoinment
	for rows.Next() {
		appointment = &appointmentEntity.Appoinment{}

		err := rows.Scan(
			&appointment.ID,
			&appointment.EmployerID,
			&appointment.CustomerID,
			&appointment.Date,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	return appointment, nil
}

func (provider *DBProvider) GetWithServicesByID(
	ctx context.Context, ID string,
) (*appointmentEntity.AppointmentWithServicesInfo, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	select_for_total_cost := fmt.Sprintf("(SELECT COALESCE(SUM(cost_of),0) FROM %s WHERE appointment_id = %s.id)", db.Service, db.Appointment)
	select_for_total_duration := fmt.Sprintf("(SELECT COALESCE(SUM(duration),0) FROM %s WHERE appointment_id = %s.id)", db.Service, db.Appointment)

	columns := fmt.Sprintf(
		"%s.id as appointment_id, %s.id as customer_id, %s.name as customer_name, %s.date_of, %s as total_cost, %s as total_duration, %s.id as service_id, %s.employee_id, %s.type_of, %s.cost_of, %s.duration, %s.is_completed, %s.observations ",
		db.Appointment, db.Customer, db.Customer, db.Appointment, select_for_total_cost, select_for_total_duration, db.Service, db.Service, db.Service, db.Service, db.Service, db.Service, db.Service,
	)

	queryCommand := fmt.Sprintf("SELECT %s FROM %s \n", columns, db.Appointment)
	queryCommand += fmt.Sprintf("INNER JOIN %s ON %s.customer_id = %s.id \n", db.Customer, db.Appointment, db.Customer)
	queryCommand += fmt.Sprintf("INNER JOIN %s ON %s.appointment_id = %s.id \n", db.Service, db.Service, db.Appointment)
	queryCommand += fmt.Sprintf("WHERE %s.id = $1", db.Appointment)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, ID)
	if err != nil {
		return nil, err
	}

	appointment := &appointmentEntity.AppointmentWithServicesInfo{}

	for rows.Next() {
		service := &serviceEntity.Service{}

		err := rows.Scan(
			&appointment.ID,
			&appointment.CustomerID,
			&appointment.CustomerName,
			&appointment.Date,
			&appointment.TotalCost,
			&appointment.TotalDuration,
			&service.ID,
			&service.EmployeeID,
			&service.Type,
			&service.Cost,
			&service.Duration,
			&service.IsCompleted,
			&service.Observations,
		)

		appointment.Services = append(appointment.Services, *service)

		if err != nil {
			return nil, err
		}
	}

	return appointment, nil
}
