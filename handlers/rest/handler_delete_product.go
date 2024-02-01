package rest

import (
	"log/slog"
	"net/http"

	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")
	if productID == "" {
		slog.Warn("field id is required")
		response.Encode(w, e.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	err := h.svcProduct.DeleteProduct(productID)
	if err != nil {
		slog.Error("failed to delete product", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, "success", http.StatusOK)
}
