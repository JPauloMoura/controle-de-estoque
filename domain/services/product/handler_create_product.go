package product

import (
	"log/slog"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
)

func (p productService) CreateProduct(product entity.Product) (*entity.Product, error) {
	productInserted, err := p.database.InsertProduct(product)
	if err != nil {
		slog.Error("failed to create product", err)
		return nil, err
	}

	return productInserted, nil
}
