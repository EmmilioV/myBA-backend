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

type DBProvider struct {
	DBConnection *database.DBConnection
}

func NewDBProvider(
	dbConnection *database.DBConnection,
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

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", CUSTOMER_TABLE_NAME)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, customerID)
	if err != nil {
		return nil, err
	}

	customer := entity.Customer{}
	for rows.Next() {
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

	return &customer, nil
}
