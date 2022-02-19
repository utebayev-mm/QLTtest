package utils

import (
	"database/sql"
)

type Database struct {
	SQLDB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{SQLDB: db}
}
func ConnectDb(driverName string, SqlDbName string) (*Database, error) {
	SQLDB, err := sql.Open(driverName, SqlDbName)
	if err != nil {
		return nil, err
	}
	if err = SQLDB.Ping(); err != nil {
		return nil, err
	}
	return &Database{SQLDB}, nil
}
