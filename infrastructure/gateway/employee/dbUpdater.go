package employee

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
	db "go.mod/infrastructure/database"
)

type DBUpdater struct {
	DBConnection *db.Connection
}

func NewDBUpdater(
	dbConnection *db.Connection,
) employeeGateway.IDBUpdater {
	return &DBUpdater{
		DBConnection: dbConnection,
	}
}

func (updater *DBUpdater) UpdateByID(
	ctx context.Context, employee *entity.Employee,
) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	updateCommand := fmt.Sprintf("UPDATE name = $1, email = $2 SET %s WHERE id=$3", db.Employee)

	result, err := updater.DBConnection.SQL_DB.ExecContext(ctxTimeout, updateCommand, employee.Name, employee.Email, employee.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected <= 0 {
		return errors.New("EMPLOYEE_NOT_FOUND")
	}

	log.Printf("INFO: rows affected updating employee: %d", rowsAffected)

	return nil
}
