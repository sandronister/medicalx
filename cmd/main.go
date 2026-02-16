package main

import (
	"log"
	"os"

	"github.com/sandronister/medicalx/internal/infra/database/connection"
	"github.com/sandronister/medicalx/internal/infra/web/server"
)

func main() {
	dbType := connection.DBType(getEnv("DB_TYPE", string(connection.PostgreSQL)))

	db, err := connection.NewConnection(dbType)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	addr := ":" + getEnv("PORT", "3000")

	srv := server.NewServer(db)
	if err := srv.Start(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
