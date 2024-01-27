package product

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/14-web_api/services/product/entity"
)

func (p productService) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Warn("invalid method",
			slog.String("aceppted", http.MethodPost),
			slog.String("received", r.Method),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid method")
		return
	}

	idInt, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		slog.Error("failed to convert product id to int", err)
		return
	}

	priceFloat, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		slog.Error("failed to convert product price", err)
		return
	}

	availableQuantityInt, err := strconv.Atoi(r.FormValue("availableQuantity"))
	if err != nil {
		slog.Error("failed to convert product availableQuantity", err)
		return
	}

	product := entity.Product{
		Id:                idInt,
		Name:              r.FormValue("name"),
		Price:             priceFloat,
		Description:       r.FormValue("description"),
		AvailableQuantity: availableQuantityInt,
	}

	p.database.UpdateProduct(product)

	http.Redirect(w, r, "/", 301)
}
