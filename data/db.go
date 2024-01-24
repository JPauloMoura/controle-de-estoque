package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connect, err := sql.Open("postgres", getConnectStr())
	if err != nil || connect == nil {
		log.Fatal("failed to open conection: ", err)
	}

	if err := connect.Ping(); err != nil {
		log.Fatalf("failed to ping on database: %v\nconnectionString: %s\n", err, getConnectStr())
	}
	return connect
}

func getConnectStr() string {
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")
	// postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full

	// return fmt.Sprintf("postgres://pqgotest:password@localhost/pqgotest?sslmode=%s", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_SSLMODE)
	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_SSLMODE)
}
