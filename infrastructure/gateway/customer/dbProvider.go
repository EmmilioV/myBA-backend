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

type DBProvider struct {
	DBConnection *db.Connection
}

func NewDBProvider(
	dbConnection *db.Connection,
) customerGateway.IDBProvider {
	return &DBProvider{
		DBConnection: dbConnection,
	}
}

func (provider *DBProvider) GetByID(
	ctx context.Context, customerID string,
) (*entity.Customer, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", db.Customer)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, customerID)
	if err != nil {
		return nil, err
	}

	var customer *entity.Customer

	for rows.Next() {
		customer = &entity.Customer{}

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.PhoneNumber,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	return customer, nil
}
