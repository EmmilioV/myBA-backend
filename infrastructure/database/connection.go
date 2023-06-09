package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Connection struct {
	SQL_DB *sql.DB
}

func NewConnection(
	dbSettings *Settings,
) *Connection {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbSettings.Host, dbSettings.Port, dbSettings.User, dbSettings.Password,
		dbSettings.Name, dbSettings.SslMode,
	)

	sql_db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return &Connection{
		SQL_DB: sql_db,
	}
}
