package employee

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
	"go.mod/infrastructure/database"
)

type DBInserter struct {
	DBConnection *database.DBConnection
}

func NewDBInserter(
	dbConnection *database.DBConnection,
) employeeGateway.IDBInserter {
	return &DBInserter{
		DBConnection: dbConnection,
	}
}

func (inserter *DBInserter) InsertOne(
	ctx context.Context, employee *entity.Employee,
) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	values := fmt.Sprintf("'%s','%s','%s'", employee.ID, employee.Name, employee.Email)
	insertCommand := fmt.Sprintf("INSERT INTO %s (id, name, email) VALUES (%s)", EMPLOYEE_TABLE_NAME, values)

	result, err := inserter.DBConnection.SQL_DB.ExecContext(ctxTimeout, insertCommand)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("INFO: rows affected inserting employee: %d", rowsAffected)

	return nil
}
