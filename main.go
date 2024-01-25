package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/14-web_api/pkg/loggers"
	"github.com/14-web_api/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("failed to loading .env file")
	}

	loggers.ConfigLogger()
	routes.Handler()

	slog.Info("server is running in port 3002...")
	http.ListenAndServe(":3002", nil)
}
