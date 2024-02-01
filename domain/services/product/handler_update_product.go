package product

import (
	"log/slog"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
)

func (p productService) UpdateProduct(product entity.Product) error {
	if err := p.database.UpdateProduct(product); err != nil {
		slog.Error("failed tp updade product", err)
		return err
	}

	return nil
}
