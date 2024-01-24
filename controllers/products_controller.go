package controllers

import (
	"html/template"
	"log"
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
	if r.Method == "POST" {
		priceFloat, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("InsertProduct: conversão de preço", err)
		}

		availableQuantityInt, err := strconv.Atoi(r.FormValue("availableQuantity"))
		if err != nil {
			log.Println("InsertProduct: conversão de availableQuantity", err)
		}

		p := models.Product{
			Name:              r.FormValue("name"),
			Description:       r.FormValue("description"),
			Price:             priceFloat,
			AvailableQuantity: availableQuantityInt,
		}

		models.InsertProduct(p)
	}

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
	if r.Method == "POST" {
		idInt, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("UpdateProduct id:", err.Error())
		}

		priceFloat, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("UpdateProduct price:", err.Error())
		}

		availableQuantityInt, err := strconv.Atoi(r.FormValue("availableQuantity"))
		if err != nil {
			log.Println("UpdateProduct availableQuantity:", err.Error())
		}

		p := models.Product{
			Id:                idInt,
			Name:              r.FormValue("name"),
			Price:             priceFloat,
			Description:       r.FormValue("description"),
			AvailableQuantity: availableQuantityInt,
		}

		models.UpdateProduct(p)
	}
	http.Redirect(w, r, "/", 301)
}
