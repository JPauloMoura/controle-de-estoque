package rest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
	"github.com/go-chi/chi/v5"
)

func (h handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		slog.Warn("invalid method",
			slog.String("aceppted", http.MethodPut),
			slog.String("received", r.Method),
		)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid method")
		return
	}

	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to convert product id", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("field id invalid"))
		return
	}

	var product entity.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		slog.Error("failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	product.Id = productID
	err = h.svcProduct.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("no updated"))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("sucess")
}
