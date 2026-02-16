package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "medicalx")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		user, password, host, port, dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening MySQL database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to MySQL database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return db, nil
}
