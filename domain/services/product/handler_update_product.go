package product

import "github.com/JPauloMoura/controle-de-estoque/domain/entity"

func (p productService) UpdateProduct(product entity.Product) error {
	p.database.UpdateProduct(product)

	return nil
}
