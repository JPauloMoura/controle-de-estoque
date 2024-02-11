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
	"github.com/JPauloMoura/controle-de-estoque/pkg/configs"
	"github.com/JPauloMoura/controle-de-estoque/pkg/loggers"
	"github.com/JPauloMoura/controle-de-estoque/pkg/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("failed to loading .env file")
	}

	cfg := configs.BuildConfig()

	loggers.ConfigLogger()
	handlers := buildHandlers(cfg)

	slog.Info("server is running in port " + cfg.ServerPort())
	http.ListenAndServe(":"+cfg.ServerPort(), handlers)
}

func buildHandlers(cfg *configs.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middlewares.Cors())

	db := database.ConnectDb(cfg)
	productRepository := repository.NewProductRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	productService := product.NewProductService(productRepository)
	authorization := auth.NewJwtAuth(cfg.JwtKey())

	{ // user handlers
		h := userHandler.NewHandlerUser(userRepository, authorization)
		router.Post("/user", h.CreateUser)
		router.Post("/user/login", h.Login)
	}

	{ // product handlers
		h := rest.NewHandlerProduct(productService)
		router.Group(func(r chi.Router) {
			r.Use(authorization.MiddlewareAuth)

			r.Get("/products", h.ListProducts)
			r.Get("/products/{id}", h.GetProduct)
			r.Post("/products", h.CreateProduct)
			r.Put("/products/{id}", h.UpdateProduct)
			r.Delete("/products/{id}", h.DeleteProduct)

		})
	}

	return router
}
