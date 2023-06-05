package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/service/entity"
	"go.mod/domain/service/gateway"
	db "go.mod/infrastructure/database"
)

type DBInserter struct {
	DBConnection *db.DBConnection
}

func NewDBInserter(
	dbConnection *db.DBConnection,
) gateway.IDBInserter {
	return &DBInserter{
		DBConnection: dbConnection,
	}
}

func (inserter *DBInserter) InsertOne(
	ctx context.Context, service *entity.Service,
) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertCommand := fmt.Sprintf(
		"INSERT INTO %s (appointment_id, employee_id, type_of, cost_of, duration, is_completed, observations) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		db.Service,
	)

	result, err := inserter.DBConnection.SQL_DB.ExecContext(ctxTimeout, insertCommand,
		service.AppointmentID,
		service.EmployeeID,
		service.Type,
		service.Cost,
		service.Duration,
		service.IsCompleted,
		service.Observations,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("INFO: rows affected inserting service: %d", rowsAffected)

	return nil
}
