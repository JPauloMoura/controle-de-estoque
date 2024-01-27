package product

import (
	"html/template"
	"net/http"

	"github.com/14-web_api/repository"
)

type ProductService interface {
	HandleIndex(w http.ResponseWriter, r *http.Request)
	HandlerForm(w http.ResponseWriter, r *http.Request)
	InsertProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	EditProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	//TODO: GET product, List product

}

func NewProductService(repo repository.ProductRepository) ProductService {
	var temp = template.Must(template.ParseGlob("templates/*.html"))

	return productService{
		database:   repo,
		renderTmpl: temp,
	}
}

type productService struct {
	database   repository.ProductRepository
	renderTmpl *template.Template
}

// Esse methodo deve ser movido para uma outra estrutura, responsavel pela view
func (p productService) HandlerForm(w http.ResponseWriter, r *http.Request) {
	p.renderTmpl.ExecuteTemplate(w, "Form", nil)
}
