package product

import (
	"log/slog"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
)

func (p productService) GetProduct(productID int) (*entity.Product, error) {
	product, err := p.database.GetProduct(productID)
	if err != nil {
		slog.Error("failed to get product", err)
		return nil, err
	}

	return product, nil
}
