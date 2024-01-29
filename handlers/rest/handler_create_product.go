package rest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
)

func (h handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Warn("invalid method",
			slog.String("aceppted", http.MethodPost),
			slog.String("received", r.Method),
		)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("invalid method"))
		return
	}

	var product entity.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		slog.Error("failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := product.Validate(); err != nil {
		slog.Error("failed to validate product", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := h.svcProduct.CreateProduct(product); err != nil {
		slog.Error("failed to create product", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("internal server error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("sucess")
}
