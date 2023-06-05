package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/service/entity"
	serviceGateway "go.mod/domain/service/gateway"
	"go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *database.DBConnection
}

func NewDBProvider(
	dbConnection *database.DBConnection,
) serviceGateway.IDBProvider {
	return &DBProvider{
		DBConnection: dbConnection,
	}
}

func (provider *DBProvider) GetByID(
	ctx context.Context, serviceID string,
) (*entity.Service, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", SERVICE_TABLE_NAME)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, serviceID)
	if err != nil {
		return nil, err
	}

	var service *entity.Service
	for rows.Next() {
		service = &entity.Service{}

		err := rows.Scan(
			&service.ID,
			&service.AppointmentID,
			&service.EmployeeID,
			&service.Type,
			&service.Cost,
			&service.Duration,
			&service.IsCompleted,
			&service.Observations,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	return service, nil
}

func (provider *DBProvider) GetAll(
	ctx context.Context,
) ([]*entity.Service, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s", SERVICE_TABLE_NAME)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand)
	if err != nil {
		return nil, err
	}

	var services []*entity.Service
	for rows.Next() {
		service := &entity.Service{}

		err := rows.Scan(
			&service.ID,
			&service.AppointmentID,
			&service.EmployeeID,
			&service.Type,
			&service.Cost,
			&service.Duration,
			&service.IsCompleted,
			&service.Observations,
		)

		services = append(services, service)

		if err != nil {
			log.Fatal(err)
		}
	}

	return services, nil
}
