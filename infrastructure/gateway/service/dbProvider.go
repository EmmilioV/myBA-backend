package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/service/entity"
	serviceGateway "go.mod/domain/service/gateway"
	db "go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *db.DBConnection
}

func NewDBProvider(
	dbConnection *db.DBConnection,
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

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", db.Service)

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

	queryCommand := fmt.Sprintf("SELECT * FROM %s", db.Service)

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
