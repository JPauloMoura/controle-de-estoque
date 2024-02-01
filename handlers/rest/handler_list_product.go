package rest

import (
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
)

func (h handlerProduct) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svcProduct.ListProducts()
	if err != nil {
		slog.Error("failed to list products", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, products, http.StatusOK)
}
