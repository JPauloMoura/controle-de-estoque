package database

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/JPauloMoura/controle-de-estoque/pkg/configs"
	_ "github.com/lib/pq"
)

func ConnectDb(cfg *configs.Config) *sql.DB {
	connectionString := cfg.DbConnectionStr()

	connect, err := sql.Open("postgres", connectionString)
	if err != nil || connect == nil {
		slog.Debug("failed to open conection. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed to open conection", err)
	}

	if err := connect.Ping(); err != nil {
		slog.Debug("failed to ping on database. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed to ping on database", err)
	}
	return connect
}
