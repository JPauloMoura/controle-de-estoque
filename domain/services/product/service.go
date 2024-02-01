package product

import (
	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
	"github.com/JPauloMoura/controle-de-estoque/domain/repository"
)

type ProductService interface {
	ListProducts() ([]entity.Product, error)
	CreateProduct(entity.Product) (*entity.Product, error)
	DeleteProduct(productID int) error
	GetProduct(productID int) (*entity.Product, error)
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
