package rest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func (h handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	if productID == "" {
		slog.Warn("field id is required")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("invalid id"))
		return
	}

	h.svcProduct.GetProduct(productID)

	product, err := h.svcProduct.GetProduct(productID)
	if err != nil {
		slog.Error("failed to get product", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
