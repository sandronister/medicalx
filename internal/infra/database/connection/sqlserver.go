package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func NewSQLServerConnection() (*sql.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "1433")
	user := getEnv("DB_USER", "sa")
	password := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "medicalx")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, host, port, dbName,
	)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening SQL Server database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to SQL Server database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return db, nil
}
