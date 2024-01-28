package rest

import (
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/domain/services/product"
)

func Handler(svc product.ProductService) {
	h := newHandlerProduct(svc)

	http.HandleFunc("/products/create", h.CreateProduct)
	http.HandleFunc("/products/update", h.UpdateProduct)
	http.HandleFunc("/products", h.ListProducts)
	http.HandleFunc("/products/find", h.GetProduct)
	http.HandleFunc("/products/delete", h.DeleteProduct)
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
