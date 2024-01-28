package webserver

import (
	"net/http"
)

func Handler(h HandlerProduct) {
	http.HandleFunc("/", h.ListProducts)

	http.HandleFunc("/novo-produto", h.RederPageCreateProduct)
	http.HandleFunc("/insert", h.CreateProduct)
	http.HandleFunc("/delete", h.DeleteProduct)
	http.HandleFunc("/editar", h.RenderPageEditeProduct)
	http.HandleFunc("/update", h.UpdateProduct)
}
