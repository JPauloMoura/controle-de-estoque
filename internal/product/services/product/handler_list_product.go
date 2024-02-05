package product

import (
	"log/slog"

	"github.com/JPauloMoura/controle-de-estoque/internal/product/entity"
	"github.com/JPauloMoura/controle-de-estoque/internal/product/repository"
)

func (p productService) ListProducts(pagination *repository.Pagination) ([]entity.Product, error) {
	listProduct, err := p.database.GetAllProducts(pagination)
	if err != nil {
		slog.Error("failed to list products", err)
		return nil, err
	}

	return listProduct, nil
}
