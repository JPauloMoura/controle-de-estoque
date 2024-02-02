package rest

import (
	"log/slog"
	"net/http"
	"strconv"

	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to convert product id", err)
		response.Encode(w, e.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	product, err := h.svcProduct.GetProduct(productID)
	if err != nil {
		slog.Error("failed to get product", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, product, http.StatusOK)
}
