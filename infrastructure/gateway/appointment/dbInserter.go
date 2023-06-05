package appointment

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/appointment/entity"
	appintmentGateway "go.mod/domain/appointment/gateway"
	"go.mod/infrastructure/database"
)

type DBInserter struct {
	DBConnection *database.DBConnection
}

func NewDBInserter(
	dbConnection *database.DBConnection,
) appintmentGateway.IDBInserter {
	return &DBInserter{
		DBConnection: dbConnection,
	}
}

func (inserter *DBInserter) InsertOne(
	ctx context.Context, appointment *entity.Appoinment,
) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertCommand := fmt.Sprintf("INSERT INTO %s (employer_id, customer_id, date_of) VALUES ($1, $2, $3)", APPOINTMENT_TABLE_NAME)

	result, err := inserter.DBConnection.SQL_DB.ExecContext(ctxTimeout, insertCommand, appointment.EmployerID, appointment.CustomerID, appointment.Date)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("INFO: rows affected inserting appointment: %d", rowsAffected)

	return nil
}
