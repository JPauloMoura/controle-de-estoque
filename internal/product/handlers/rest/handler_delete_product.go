package rest

import (
	"log/slog"
	"net/http"
	"strconv"

	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to convert product id", err)
		response.Encode(w, e.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	if err := h.svcProduct.DeleteProduct(productID); err != nil {
		slog.Error("failed to delete product", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, "success", http.StatusOK)
}
