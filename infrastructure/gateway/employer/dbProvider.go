package employer

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mod/domain/employer/entity"
	employerGateway "go.mod/domain/employer/gateway"
	"go.mod/infrastructure/database"
)

type DBProvider struct {
	DBConnection *database.DBConnection
}

func NewDBProvider(
	dbConnection *database.DBConnection,
) employerGateway.IDBProvider {
	return &DBProvider{
		dbConnection,
	}
}

func (provider *DBProvider) GetByID(
	ctx context.Context, employerID string,
) (*entity.Employer, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	filter := fmt.Sprintf("id='%s'", employerID)

	queryCommand := fmt.Sprintf("SELECT * FROM %s WHERE %s", TABLE_NAME, filter)

	rows, err := provider.DBConnection.SQL_DB.QueryContext(ctxTimeout, queryCommand)
	if err != nil {
		return nil, err
	}

	employer := entity.Employer{}
	for rows.Next() {
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

	return &employer, nil
}