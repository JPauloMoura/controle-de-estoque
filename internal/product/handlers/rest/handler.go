package rest

import (
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/product/services/product"
)

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

func NewHandlerProduct(svcProduct product.ProductService) HandlerProduct {
	return handlerProduct{
		svcProduct: svcProduct,
	}
}
