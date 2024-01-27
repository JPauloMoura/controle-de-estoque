package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/14-web_api/handlers"
	"github.com/14-web_api/infrastructure"
	"github.com/14-web_api/pkg/loggers"
	"github.com/14-web_api/repository"
	"github.com/14-web_api/services/product"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("failed to loading .env file")
	}

	loggers.ConfigLogger()
	buildDependencies()

	slog.Info("server is running in port 3002...")
	http.ListenAndServe(":3002", nil)
}

func buildDependencies() {
	db := infrastructure.ConnectDb()
	repo := repository.NewProductRepository(db)
	svc := product.NewProductService(repo)
	handlers.Handler(svc)
}
