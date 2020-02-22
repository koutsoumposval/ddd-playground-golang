package config

import (
	"database/sql"
	"fmt"

	"github.com/koutsoumposval/ddd-playground-golang/app/environment"
)

// DatabaseConnection is used for database connection.
func DatabaseConnection() (*sql.DB, error) {
	user := environment.GetEnvVar("DB_USER", "user")
	password := environment.GetEnvVar("DB_PASSWORD", "p@ssw0rd")
	host := environment.GetEnvVar("DB_HOST", "localhost")
	port := environment.GetEnvVar("DB_PORT", "33061")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ddd_playground_golang?parseTime=true", user, password, host, port)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
