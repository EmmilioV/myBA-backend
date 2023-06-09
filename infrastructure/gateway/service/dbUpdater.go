package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mod/domain/service/entity"
	serviceGateway "go.mod/domain/service/gateway"
	db "go.mod/infrastructure/database"
)

type DBUpdater struct {
	DBConnection *db.Connection
}

func NewDBUpdater(
	dbConnection *db.Connection,
) serviceGateway.IDBUpdater {
	return &DBUpdater{
		DBConnection: dbConnection,
	}
}

func (updater *DBUpdater) UpdateOneByID(
	ctx context.Context, service *entity.Service,
) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	updateCommand := fmt.Sprintf("UPDATE %s SET duration = $1, is_completed = $2, observations = $3 WHERE id = $4", db.Service)

	result, err := updater.DBConnection.SQL_DB.ExecContext(ctxTimeout, updateCommand, service.Duration, service.IsCompleted, service.Observations, service.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected <= 0 {
		return errors.New("SERVICE_NOT_FOUND")
	}

	log.Printf("INFO: rows affected updating service: %d", rowsAffected)

	return nil
}
