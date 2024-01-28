package product

import (
	"github.com/14-web_api/domain/entity"
)

func (p productService) ListProducts() ([]entity.Product, error) {
	listProduct := p.database.GetAllProducts()
	return listProduct, nil
}
