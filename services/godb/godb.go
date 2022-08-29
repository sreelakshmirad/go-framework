package godb

import (
	"database/sql"
)

// interface for all database types
type DBConnector interface {
	CreateDBConnection() (*sql.DB, error)
}

type ClientDBConfig struct {
	DBSecret string `json:"secret"`
	DBDriver string `json:"driver"`
	DBName   string `json:"dbname"`
}

// func NewDbConnector accepts a string to specify the driver type and returns a DB struct
func NewDbConnector(driver string) (castType DBConnector) {
	switch driver {
	case "postgres":
		castType = GoDbPostgres{}
	case "mysql":
		castType = GoDbMysql{}
	default:
		castType = GoDbPostgres{}

	}
	return
}
