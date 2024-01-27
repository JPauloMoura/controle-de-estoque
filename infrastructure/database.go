package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connect, err := sql.Open("postgres", getConnectStr())
	if err != nil || connect == nil {
		slog.Error("failed to open conection", err, slog.String("connectionString", getConnectStr()))
		log.Fatal("down service")
	}

	if err := connect.Ping(); err != nil {
		slog.Error("failed to ping on database", err, slog.String("connectionString", getConnectStr()))
		log.Fatal("down service")
	}
	return connect
}

func getConnectStr() string {
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_SSLMODE)
}
