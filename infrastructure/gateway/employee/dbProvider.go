package employee

import (
	"context"
	"fmt"
	"log"
	"time"

	employeeEntity "go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
	serviceEntity "go.mod/domain/service/entity"
	db "go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *db.DBConnection
}

func NewDBProvider(
	dbConnection *db.DBConnection,
) employeeGateway.IDBProvider {
	return &DBProvider{
		DBConnection: dbConnection,
	}
}

func (provider *DBProvider) GetByID(
	ctx context.Context, employeeID string,
) (*employeeEntity.Employee, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", db.Employee)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, employeeID)
	if err != nil {
		return nil, err
	}

	var employee *employeeEntity.Employee
	for rows.Next() {
		employee = &employeeEntity.Employee{}

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	return employee, nil
}

func (provider *DBProvider) GetByIDWithServicesInformation(
	ctx context.Context, employeeID string,
) (*employeeEntity.EmployeeWithServicesInfo, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	columns := fmt.Sprintf(
		"%s.id, %s.name, %s.email, %s.date_of, %s.id as service_id, %s.type_of, %s.cost_of, %s.duration, %s.is_completed, %s.observations, %s.name as customer_name",
		db.Employee, db.Employee, db.Employee, db.Appointment, db.Service, db.Service, db.Service, db.Service, db.Service, db.Service, db.Customer,
	)

	queryCommand := fmt.Sprintf("SELECT %s FROM %s \n", columns, db.Employee)
	queryCommand += fmt.Sprintf("INNER JOIN %s ON %s.employee_id = %s.id \n", db.Service, db.Service, db.Employee)
	queryCommand += fmt.Sprintf("INNER JOIN %s ON %s.id = %s.appointment_id \n", db.Appointment, db.Appointment, db.Service)
	queryCommand += fmt.Sprintf("INNER JOIN %s ON %s.id = %s.customer_id \n", db.Customer, db.Customer, db.Appointment)
	queryCommand += fmt.Sprintf("WHERE %s.id = $1", db.Employee)
	queryCommand += fmt.Sprintf("ORDER BY %s.date_of", db.Appointment)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, employeeID)
	if err != nil {
		return nil, err
	}

	employee := &employeeEntity.EmployeeWithServicesInfo{}

	for rows.Next() {
		service := &serviceEntity.ServiceWithAdditionalInfo{}

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&service.Date,
			&service.ID,
			&service.Type,
			&service.Cost,
			&service.Duration,
			&service.IsCompleted,
			&service.Observations,
			&service.CustomerName,
		)

		employee.Services = append(employee.Services, service)

		if err != nil {
			log.Fatal(err)
		}
	}

	return employee, nil
}
