package employee

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
	"go.mod/infrastructure/database"
)

type DBUpdater struct {
	DBConnection *database.DBConnection
}

func NewDBUpdater(
	dbConnection *database.DBConnection,
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

	set := fmt.Sprintf("name = '%s', email = '%s'", employee.Name, employee.Email)
	condition := fmt.Sprintf("id='%s'", employee.ID)
	updateCommand := fmt.Sprintf("UPDATE %s SET %s WHERE %s", EMPLOYEE_TABLE_NAME, set, condition)

	result, err := updater.DBConnection.SQL_DB.ExecContext(ctxTimeout, updateCommand)
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
