package rest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")
	if productID == "" {
		slog.Warn("field id is required")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("invalid id"))
		return
	}

	err := h.svcProduct.DeleteProduct(productID)
	if err != nil {
		slog.Error("failed to delete product", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("sucess")
}
