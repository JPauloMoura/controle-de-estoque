package product

import "github.com/JPauloMoura/controle-de-estoque/domain/entity"

func (p productService) CreateProduct(product entity.Product) error {
	p.database.InsertProduct(product)
	return nil
}
