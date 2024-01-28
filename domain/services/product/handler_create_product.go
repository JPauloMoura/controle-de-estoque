package product

import "github.com/14-web_api/domain/entity"

func (p productService) CreateProduct(product entity.Product) error {
	p.database.InsertProduct(product)
	return nil
}
