package product

import (
	"github.com/14-web_api/domain/entity"
	"github.com/14-web_api/domain/repository"
)

type ProductService interface {
	ListProducts() ([]entity.Product, error)
	CreateProduct(entity.Product) error
	DeleteProduct(productID string) error
	GetProduct(productID string) (entity.Product, error)
	UpdateProduct(product entity.Product) error
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return productService{
		database: repo,
	}
}

type productService struct {
	database repository.ProductRepository
}
