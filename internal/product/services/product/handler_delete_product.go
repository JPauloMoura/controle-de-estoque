package product

import "log/slog"

func (p productService) DeleteProduct(productID int) error {
	if err := p.database.DeleteProduct(productID); err != nil {
		slog.Error("failed to delete product", err)
		return err
	}

	return nil
}
