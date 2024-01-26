package controllers

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/14-web_api/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	listProduct := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", listProduct)
}

func HandlerForm(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Form", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
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

	p := models.Product{
		Name:              r.FormValue("name"),
		Description:       r.FormValue("description"),
		Price:             priceFloat,
		AvailableQuantity: availableQuantityInt,
	}

	models.InsertProduct(p)

	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	models.DeleteProduct(productID)

	http.Redirect(w, r, "/", 301)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	product := models.GetProduct(productID)
	temp.ExecuteTemplate(w, "Edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	p := models.Product{
		Id:                idInt,
		Name:              r.FormValue("name"),
		Price:             priceFloat,
		Description:       r.FormValue("description"),
		AvailableQuantity: availableQuantityInt,
	}

	models.UpdateProduct(p)

	http.Redirect(w, r, "/", 301)
}
