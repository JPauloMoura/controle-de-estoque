package webserver

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"

	"github.com/JPauloMoura/controle-de-estoque/internal/product/entity"
	"github.com/JPauloMoura/controle-de-estoque/internal/product/services/product"
	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
)

type HandlerProduct interface {
	ListProducts(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)

	RenderPageEditeProduct(w http.ResponseWriter, r *http.Request)
	RederPageCreateProduct(w http.ResponseWriter, r *http.Request)
}

func NewHandlerProduct(svcProduct product.ProductService) HandlerProduct {
	var temp = template.Must(template.ParseGlob("templates/*.html"))
	return handlerProduct{
		svcProduct: svcProduct,
		templates:  temp,
	}
}

type handlerProduct struct {
	svcProduct product.ProductService
	templates  *template.Template
}

func (h handlerProduct) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svcProduct.ListProducts(nil)
	if err != nil {
		slog.Error("failed to list products")
		return
	}

	err = h.templates.ExecuteTemplate(w, "Index", products)
	if err != nil {
		slog.Error("failed to execute Index template", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("page unavailable"))
		return
	}
}

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

	priceFloat, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		slog.Error("failed to convert product price", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("field price invalid"))
		return
	}

	availableQuantityInt, err := strconv.Atoi(r.FormValue("availableQuantity"))
	if err != nil {
		slog.Error("failed to convert product availableQuantity", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("field availableQuantity invalid"))
		return
	}

	product := entity.Product{
		Name:              r.FormValue("name"),
		Description:       r.FormValue("description"),
		Price:             priceFloat,
		AvailableQuantity: availableQuantityInt,
	}

	_, err = h.svcProduct.CreateProduct(product)
	if err != nil {
		slog.Error("failed to create product", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("internal server error"))
		return
	}

	http.Redirect(w, r, "/", 301)
}

func (h handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("failed to convert product id", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e.ErrorInvalidId)
		return
	}

	err = h.svcProduct.DeleteProduct(productID)
	if err != nil {
		slog.Error("failed to delete product", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	http.Redirect(w, r, "/", 301)
}

func (h handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Warn("invalid method",
			slog.String("aceppted", http.MethodPost),
			slog.String("received", r.Method),
		)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid method")
		return
	}

	idInt, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		slog.Error("failed to convert product id", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("field id invalid"))
		return
	}

	priceFloat, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		slog.Error("failed to convert product price", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("field price invalid"))
		return
	}

	availableQuantityInt, err := strconv.Atoi(r.FormValue("availableQuantity"))
	if err != nil {
		slog.Error("failed to convert product availableQuantity", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("field availableQuantity invalid"))
		return
	}

	//Como é atualização, tudo bem alguns dados serem vazios
	product := entity.Product{
		Id:                idInt,
		Name:              r.FormValue("name"),
		Price:             priceFloat,
		Description:       r.FormValue("description"),
		AvailableQuantity: availableQuantityInt,
	}

	h.svcProduct.UpdateProduct(product)

	http.Redirect(w, r, "/", 301)
}

func (h handlerProduct) RenderPageEditeProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("failed to convert product id", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e.ErrorInvalidId)
		return
	}

	product, err := h.svcProduct.GetProduct(productID)
	if err != nil {
		slog.Error("failed to get product", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	h.templates.ExecuteTemplate(w, "Edit", product)
}

func (h handlerProduct) RederPageCreateProduct(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "Form", nil)
}
