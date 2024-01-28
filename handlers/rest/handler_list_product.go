package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h handlerProduct) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svcProduct.ListProducts()
	if err != nil {
		slog.Error("failed to list products")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
