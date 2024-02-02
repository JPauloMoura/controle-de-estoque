package rest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/product/entity"
	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
)

func (h handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Warn("invalid method",
			slog.String("aceppted", http.MethodPost),
			slog.String("received", r.Method),
		)

		response.Encode(w, e.ErrorInvalidHttpMethod, http.StatusBadRequest)
		return
	}

	var product entity.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		slog.Error("failed to decode body", errors.Join(e.ErrorInvalidProductFieldsJson, err))
		response.Encode(w, e.ErrorInvalidProductFieldsJson, http.StatusBadRequest)
		return
	}

	if err := product.Validate(); err != nil {
		slog.Error("failed to validate product", err)
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	_, err := h.svcProduct.CreateProduct(product)
	if err != nil {
		slog.Error("failed to create product", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, "success", http.StatusCreated)
}
