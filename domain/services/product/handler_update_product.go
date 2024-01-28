package product

import "github.com/14-web_api/domain/entity"

func (p productService) UpdateProduct(product entity.Product) error {
	p.database.UpdateProduct(product)

	return nil
}
