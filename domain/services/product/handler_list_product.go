package product

import (
	"log/slog"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
)

func (p productService) ListProducts() ([]entity.Product, error) {
	listProduct, err := p.database.GetAllProducts()
	if err != nil {
		slog.Error("failed to list products", err)
		return nil, err
	}

	return listProduct, nil
}
