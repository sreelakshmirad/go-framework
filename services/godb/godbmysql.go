package godb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type GoDbMysql struct {
	Facility   string `mapstructure:"facility"`
	DBHost     string `mapstructure:"host"`
	DBPassword string `mapstructure:"password"`
	DBPort     string `mapstructure:"port"`
	DBUser     string `mapstructure:"user"`
}

func (config GoDbMysql) CreateDBConnection() (*sql.DB, error) {

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.Facility)

	db, err := sql.Open("mysql", connection)

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
