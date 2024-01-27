package product

import "net/http"

func (p productService) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	p.database.DeleteProduct(productID)

	http.Redirect(w, r, "/", 301)
}
