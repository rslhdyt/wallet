package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Connector *sql.DB

func Connect() error {
	var err error

	// TODO: get the connection string from an env variable
	Connector, err = sql.Open("mysql", "spenmo_user:spenmo@tcp(mysql:3306)/spenmo")
	// Connector, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/spenmo")

	if err != nil {
		return err
	}

	return nil
}
