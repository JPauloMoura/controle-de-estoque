package product

func (p productService) DeleteProduct(productID string) error {
	p.database.DeleteProduct(productID)
	return nil
}
