package connection

import (
	"database/sql"
	"fmt"
	"os"
)

type DBType string

const (
	MySQL      DBType = "mysql"
	PostgreSQL DBType = "postgres"
	SQLServer  DBType = "sqlserver"
)

func NewConnection(dbType DBType) (*sql.DB, error) {
	switch dbType {
	case MySQL:
		return NewMySQLConnection()
	case PostgreSQL:
		return NewPostgreSQLConnection()
	case SQLServer:
		return NewSQLServerConnection()
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
