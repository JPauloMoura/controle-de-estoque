package product

import "net/http"

func (p productService) HandleIndex(w http.ResponseWriter, r *http.Request) {
	listProduct := p.database.GetAllProducts()
	p.renderTmpl.ExecuteTemplate(w, "Index", listProduct)
}
