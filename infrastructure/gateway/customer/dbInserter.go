package customer

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/customer/entity"
	customerGateway "go.mod/domain/customer/gateway"
	"go.mod/infrastructure/database"
)

type DBInserter struct {
	DBConnection *database.DBConnection
}

func NewDBInserter(
	dbConnection *database.DBConnection,
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

	values := fmt.Sprintf("'%s','%s','%s', '%s'", customer.ID, customer.Name, customer.Email, *customer.PhoneNumber)
	insertCommand := fmt.Sprintf("INSERT INTO %s (id, name, email, phone_number) VALUES (%s)", CUSTOMER_TABLE_NAME, values)

	result, err := inserter.DBConnection.SQL_DB.ExecContext(ctxTimeout, insertCommand)
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
