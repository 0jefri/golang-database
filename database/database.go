package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnetDatabase() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=02061996 dbname=ojek_online sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
