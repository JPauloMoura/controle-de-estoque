package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/domain/repository"
	"github.com/JPauloMoura/controle-de-estoque/domain/services/product"
	"github.com/JPauloMoura/controle-de-estoque/handlers/webserver"
	"github.com/JPauloMoura/controle-de-estoque/infrastructure"
	"github.com/JPauloMoura/controle-de-estoque/pkg/loggers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("failed to loading .env file")
	}

	loggers.ConfigLogger()
	buildDependencies()

	slog.Info("webserver is running in port 8081...")
	http.ListenAndServe(":8081", nil)
}

func buildDependencies() {
	repo := repository.NewProductRepository(infrastructure.ConnectDb())
	svcProduct := product.NewProductService(repo)
	webserver.Handler(webserver.NewHandlerProduct(svcProduct))
}
