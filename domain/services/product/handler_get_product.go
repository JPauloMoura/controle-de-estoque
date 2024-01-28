package product

import "github.com/JPauloMoura/controle-de-estoque/domain/entity"

func (p productService) GetProduct(productID string) (entity.Product, error) {
	product := p.database.GetProduct(productID)
	return product, nil
}
