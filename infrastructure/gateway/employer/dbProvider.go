package employer

import (
	"context"
	"fmt"
	"log"
	"time"

	appointmentEntity "go.mod/domain/appointment/entity"
	employerEntity "go.mod/domain/employer/entity"
	employerGateway "go.mod/domain/employer/gateway"
	db "go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *db.Connection
}

func NewDBProvider(
	dbConnection *db.Connection,
) employerGateway.IDBProvider {
	return &DBProvider{
		DBConnection: dbConnection,
	}
}

func (provider *DBProvider) GetByID(
	ctx context.Context, employerID string,
) (*employerEntity.Employer, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", db.Employer)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, employerID)
	if err != nil {
		return nil, err
	}

	var employer *employerEntity.Employer
	for rows.Next() {
		employer = &employerEntity.Employer{}

		err := rows.Scan(
			&employer.ID,
			&employer.Name,
			&employer.Office,
			&employer.PhoneNumber,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	return employer, nil
}

func (provider *DBProvider) GetByIDWithAppointments(
	ctx context.Context, ID string,
) (*employerEntity.EmployerWithAppointmentsInfo, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	select_for_total_cost := fmt.Sprintf("(SELECT COALESCE(SUM(cost_of),0) FROM %s WHERE appointment_id = %s.id)", db.Service, db.Appointment)
	select_for_total_duration := fmt.Sprintf("(SELECT COALESCE(SUM(duration),0) FROM %s WHERE appointment_id = %s.id)", db.Service, db.Appointment)

	columns := fmt.Sprintf(
		"%s.id, %s.name, %s.id as appointment_id, %s.id as customer_id, %s.name as customer_name, %s.date_of, %s as total_cost, %s as total_duration",
		db.Employer, db.Employer, db.Appointment, db.Customer, db.Customer, db.Appointment, select_for_total_cost, select_for_total_duration,
	)

	queryCommand := fmt.Sprintf("SELECT %s FROM %s \n", columns, db.Employer)
	queryCommand += fmt.Sprintf("INNER JOIN %s ON %s.employer_id = %s.id \n", db.Appointment, db.Appointment, db.Employer)
	queryCommand += fmt.Sprintf("INNER JOIN %s ON %s.customer_id = %s.id \n", db.Customer, db.Appointment, db.Customer)
	queryCommand += fmt.Sprintf("WHERE %s.id = $1 \n", db.Employer)
	queryCommand += fmt.Sprintf("ORDER BY %s.date_of DESC", db.Appointment)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, ID)
	if err != nil {
		return nil, err
	}

	employer := &employerEntity.EmployerWithAppointmentsInfo{}

	for rows.Next() {
		appointment := &appointmentEntity.AppoinmentWithExtraInfo{}

		err := rows.Scan(
			&employer.ID,
			&employer.Name,
			&appointment.ID,
			&appointment.CustomerID,
			&appointment.CustomerName,
			&appointment.Date,
			&appointment.TotalCost,
			&appointment.TotalDuration,
		)

		employer.Appointments = append(employer.Appointments, appointment)

		if err != nil {
			log.Fatal(err)
		}
	}

	return employer, nil
}
