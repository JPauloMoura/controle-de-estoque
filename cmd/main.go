package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/14-web_api/domain/repository"
	"github.com/14-web_api/domain/services/product"
	"github.com/14-web_api/handlers/webserver"
	"github.com/14-web_api/infrastructure"
	"github.com/14-web_api/pkg/loggers"
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
	repo := repository.NewProductRepository(infrastructure.ConnectDb())
	svcProduct := product.NewProductService(repo)

	webserver.Handler(
		webserver.NewHandlerProduct(svcProduct),
	)

	// rest
	// handlers.HandlerRest(product.NewProductService(
	// 	repo,
	// 	rest.NewRestPresenter(),
	// ))
}
