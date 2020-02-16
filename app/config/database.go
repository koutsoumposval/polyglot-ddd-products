package config

import (
	"database/sql"
	"fmt"

	"github.com/koutsoumposval/polyglot-ddd-product/app/environment"
)

// DatabaseConnection is used for database connection.
func DatabaseConnection() (*sql.DB, error) {
	user := environment.GetEnvVar("DB_USER", "root")
	password := environment.GetEnvVar("DB_PASSWORD", "")
	host := environment.GetEnvVar("DB_HOST", "localhost")
	port := environment.GetEnvVar("DB_PORT", "3306")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/polyglot_ddd_product?parseTime=true", user, password, host, port)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
