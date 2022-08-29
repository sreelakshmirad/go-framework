package godb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type GoDbPostgres struct {
	DBName     string `mapstructure:"dbname"`
	DBUser     string `mapstructure:"user"`
	DBPassword string `mapstructure:"password"`
	DBHost     string `mapstructure:"host"`
	DBPort     string `mapstructure:"port"`
}

func (config GoDbPostgres) CreateDBConnection() (*sql.DB, error) {

	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%v",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, true)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return db, err
	}
	return db, nil

}
