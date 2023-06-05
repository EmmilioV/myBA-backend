package employee

import (
	"context"
	"fmt"
	"log"
	"time"

	employeeGateway "go.mod/domain/employee/gateway"
	db "go.mod/infrastructure/database"
)

type DBDeleter struct {
	DBConnection *db.DBConnection
}

func NewDBDeleter(
	dbConnection *db.DBConnection,
) employeeGateway.IDBDeleter {
	return &DBDeleter{
		DBConnection: dbConnection,
	}
}

func (deleter *DBDeleter) DeleteOne(
	ctx context.Context, employeeID string,
) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteCommand := fmt.Sprintf("DELETE FROM %s WHERE id=$1", db.Employee)

	result, err := deleter.DBConnection.SQL_DB.ExecContext(ctxTimeout, deleteCommand, employeeID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("INFO: rows affected deleting employee: %d", rowsAffected)

	return nil
}
