package employer

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/employer/entity"
	employerGateway "go.mod/domain/employer/gateway"
	db "go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *db.DBConnection
}

func NewDBProvider(
	dbConnection *db.DBConnection,
) employerGateway.IDBProvider {
	return &DBProvider{
		DBConnection: dbConnection,
	}
}

func (provider *DBProvider) GetByID(
	ctx context.Context, employerID string,
) (*entity.Employer, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", db.Employer)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand, employerID)
	if err != nil {
		return nil, err
	}

	var employer *entity.Employer
	for rows.Next() {
		employer = &entity.Employer{}

		err := rows.Scan(
			&employer.ID,
			&employer.Name,
			&employer.Office,
			&employer.PhoneNumber,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	return employer, nil
}
