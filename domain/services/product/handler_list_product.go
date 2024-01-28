package product

import (
	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
)

func (p productService) ListProducts() ([]entity.Product, error) {
	listProduct := p.database.GetAllProducts()
	return listProduct, nil
}
