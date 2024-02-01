package rest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		slog.Warn("invalid method",
			slog.String("aceppted", http.MethodPut),
			slog.String("received", r.Method),
		)

		response.Encode(w, e.ErrorInvalidHttpMethod, http.StatusBadRequest)
		return
	}

	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to convert product id", err)
		response.Encode(w, e.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	var product entity.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		slog.Error("failed to decode body", errors.Join(e.ErrorInvalidProductFieldsJson, err))
		response.Encode(w, e.ErrorInvalidProductFieldsJson, http.StatusBadRequest)
		return
	}

	product.Id = productID

	if err := h.svcProduct.UpdateProduct(product); err != nil {
		slog.Error("failed to update product", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, "success", http.StatusOK)
}
