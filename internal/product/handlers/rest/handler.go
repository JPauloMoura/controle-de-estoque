package rest

import (
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/product/services/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(svc product.ProductService) *chi.Mux {
	h := newHandlerProduct(svc)
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	router.Get("/products", h.ListProducts)
	router.Get("/products/{id}", h.GetProduct)
	router.Post("/products", h.CreateProduct)
	router.Put("/products/{id}", h.UpdateProduct)
	router.Delete("/products/{id}", h.DeleteProduct)

	return router
}

type HandlerProduct interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	ListProducts(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

type handlerProduct struct {
	svcProduct product.ProductService
}

func newHandlerProduct(svcProduct product.ProductService) HandlerProduct {
	return handlerProduct{
		svcProduct: svcProduct,
	}
}
