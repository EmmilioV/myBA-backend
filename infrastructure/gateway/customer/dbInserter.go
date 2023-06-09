package customer

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/customer/entity"
	customerGateway "go.mod/domain/customer/gateway"
	db "go.mod/infrastructure/database"
)

type DBInserter struct {
	DBConnection *db.Connection
}

func NewDBInserter(
	dbConnection *db.Connection,
) customerGateway.IDBInserter {
	return &DBInserter{
		DBConnection: dbConnection,
	}
}

func (inserter *DBInserter) InsertOne(
	ctx context.Context, customer *entity.Customer,
) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertCommand := fmt.Sprintf("INSERT INTO %s (id, name, email, phone_number) VALUES ($1, $2, $3, $4)", db.Customer)

	result, err := inserter.DBConnection.SQL_DB.ExecContext(ctxTimeout, insertCommand, customer.ID, customer.Name, customer.Email, *customer.PhoneNumber)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("INFO: rows affected inserting customer: %d", rowsAffected)

	return nil
}
