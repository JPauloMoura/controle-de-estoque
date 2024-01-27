package handlers

import (
	"net/http"

	"github.com/14-web_api/services/product"
)

func Handler(svc product.ProductService) {
	http.HandleFunc("/", svc.HandleIndex)
	http.HandleFunc("/novo-produto", svc.HandlerForm)
	http.HandleFunc("/insert", svc.InsertProduct)
	http.HandleFunc("/delete", svc.DeleteProduct)
	http.HandleFunc("/editar", svc.EditProduct)
	http.HandleFunc("/update", svc.UpdateProduct)
}
