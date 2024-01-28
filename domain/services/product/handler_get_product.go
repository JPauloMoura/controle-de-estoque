package product

import "github.com/14-web_api/domain/entity"

func (p productService) GetProduct(productID string) (entity.Product, error) {
	product := p.database.GetProduct(productID)
	return product, nil
}
