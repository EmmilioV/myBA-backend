package employee

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/employee/entity"
	employeeGateway "go.mod/domain/employee/gateway"
	db "go.mod/infrastructure/database"
)

type DBInserter struct {
	DBConnection *db.DBConnection
}

func NewDBInserter(
	dbConnection *db.DBConnection,
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

	insertCommand := fmt.Sprintf("INSERT INTO %s (id, name, email) VALUES ($1, $2, $3)", db.Employee)

	result, err := inserter.DBConnection.SQL_DB.ExecContext(ctxTimeout, insertCommand, employee.ID, employee.Name, employee.Email)
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
