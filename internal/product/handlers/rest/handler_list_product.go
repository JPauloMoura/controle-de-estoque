package rest

import (
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/product/repository"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
)

func (h handlerProduct) ListProducts(w http.ResponseWriter, r *http.Request) {
	pagination, err := repository.NewPagination(r)
	if err != nil {
		slog.Error("failed to create pagination", slog.Any("error", err))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	products, err := h.svcProduct.ListProducts(pagination)
	if err != nil {
		slog.Error("failed to list products", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, products, http.StatusOK)
}
