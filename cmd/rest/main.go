package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/infrastructure/database"
	"github.com/JPauloMoura/controle-de-estoque/internal/product/handlers/rest"
	"github.com/JPauloMoura/controle-de-estoque/internal/product/repository"
	"github.com/JPauloMoura/controle-de-estoque/internal/product/services/product"
	userHandler "github.com/JPauloMoura/controle-de-estoque/internal/user/handlers"
	userRepo "github.com/JPauloMoura/controle-de-estoque/internal/user/repository"
	"github.com/JPauloMoura/controle-de-estoque/pkg/auth"
	"github.com/JPauloMoura/controle-de-estoque/pkg/loggers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("failed to loading .env file")
	}

	loggers.ConfigLogger()
	handlers := buildHandlers()

	slog.Info("server is running in port 3002...")
	http.ListenAndServe(":3002", handlers)
}

func buildHandlers() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	db := database.ConnectDb()
	productRepository := repository.NewProductRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	productService := product.NewProductService(productRepository)

	{ // user handlers
		h := userHandler.NewHandlerUser(userRepository)
		router.Post("/user", h.CreateUser)
		router.Post("/user/login", h.Login)
	}

	{ // product handlers
		h := rest.NewHandlerProduct(productService)
		router.Group(func(r chi.Router) {
			r.Use(auth.MiddlewareAuth)

			r.Get("/products", h.ListProducts)
			r.Get("/products/{id}", h.GetProduct)
			r.Post("/products", h.CreateProduct)
			r.Put("/products/{id}", h.UpdateProduct)
			r.Delete("/products/{id}", h.DeleteProduct)

		})
	}

	return router
}
