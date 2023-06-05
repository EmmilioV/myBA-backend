package appointment

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/appointment/entity"
	appointmentGateway "go.mod/domain/appointment/gateway"
	"go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *database.DBConnection
}

func NewDBProvider(
	dbConnection *database.DBConnection,
) appointmentGateway.IDBProvider {
	return &DBProvider{
		DBConnection: dbConnection,
	}
}

func (provider *DBProvider) GetOneByCustomerIDAndDate(
	ctx context.Context, customerID string, date time.Time,
) (*entity.Appoinment, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE customer_id=$1 AND date_of=$2", APPOINTMENT_TABLE_NAME)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, customerID, date)
	if err != nil {
		return nil, err
	}

	var appointment *entity.Appoinment
	for rows.Next() {
		appointment = &entity.Appoinment{}

		err := rows.Scan(
			&appointment.ID,
			&appointment.EmployerID,
			&appointment.CustomerID,
			&appointment.Date,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	return appointment, nil
}
