package product

import "net/http"

func (p productService) EditProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	product := p.database.GetProduct(productID)
	p.renderTmpl.ExecuteTemplate(w, "Edit", product)
}
