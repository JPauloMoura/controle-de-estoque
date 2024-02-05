package repository

import (
	"database/sql"
	"log/slog"

	"github.com/JPauloMoura/controle-de-estoque/internal/product/entity"
	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
)

type ProductRepository interface {
	GetAllProducts(pagination *Pagination) ([]entity.Product, error)
	InsertProduct(p entity.Product) (*entity.Product, error)
	DeleteProduct(id int) error
	GetProduct(id int) (*entity.Product, error)
	UpdateProduct(p entity.Product) error
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r repository) GetAllProducts(pagination *Pagination) ([]entity.Product, error) {
	items, err := r.db.Query(pagination.Query())
	if err != nil {
		slog.Error("failed to get products: ", err)
		return nil, e.ErrorUnableToListProducts
	}

	listProduct := []entity.Product{}

	for items.Next() {
		var id, availableQuantity int
		var name, description string
		var price float64

		// its parameters must be in the same order as the entity fields

		if err := items.Scan(&id, &name, &price, &description, &availableQuantity); err != nil {
			slog.Error("failed to scan products: ", err)
			return nil, e.ErrorUnableToScanProduct
		}

		p := entity.Product{Id: id, Name: name, Price: price, Description: description, AvailableQuantity: availableQuantity}

		listProduct = append(listProduct, p)
	}

	return listProduct, nil
}

func (r repository) InsertProduct(p entity.Product) (*entity.Product, error) {
	queryInsert, err := r.db.Prepare(`
		insert into products(name, description, price, available_quantity)
		values($1, $2, $3, $4) RETURNING id
	`)

	if err != nil {
		slog.Error("failed to prepare query", err)
		return nil, e.ErrorQueryToInsertProductIsInvalid
	}

	entityDb, err := queryInsert.Exec(p.Name, p.Description, p.Price, p.AvailableQuantity)
	if err != nil {
		slog.Error("failed to insert product", err)
		return nil, e.ErrorUnableToInsertProduct
	}

	id, err := entityDb.LastInsertId()
	if err != nil {
		slog.Warn("failed get last insert id", slog.String("error", err.Error()))
		return &p, nil
	}

	p.Id = int(id)
	return &p, nil
}

func (r repository) DeleteProduct(id int) error {
	query, err := r.db.Prepare(`DELETE FROM products WHERE id=$1`)
	if err != nil {
		slog.Error("failed to prepare query to delete product", err)
		return e.ErrorQueryToDeleteProductIsInvalid
	}

	if _, err = query.Exec(id); err != nil {
		slog.Error("failed to delete product", err)
		return e.ErrorUnableToDeleteProduct
	}

	return nil
}

func (r repository) GetProduct(id int) (*entity.Product, error) {
	items, err := r.db.Query(`SELECT * FROM products WHERE id=$1`, id)
	if err != nil {
		slog.Error("failed to get product by id", err, slog.Int("productId", id))
		return nil, e.ErrorUnableToGetProducts
	}

	var p entity.Product

	if !items.Next() {
		return nil, e.ErrorProductNotFound
	}

	// its parameters must be in the same order as the entity fields
	if err := items.Scan(&p.Id, &p.Name, &p.Price, &p.Description, &p.AvailableQuantity); err != nil {
		slog.Error("failed to scan when get product by id", err, slog.Int("productId", id))
		return nil, e.ErrorUnableToScanProduct
	}

	return &p, nil
}

func (r repository) UpdateProduct(p entity.Product) error {
	query, err := r.db.Prepare(`
		UPDATE products SET 
		name=$1, price=$2, description=$3, available_quantity=$4
		WHERE id=$5
	`)

	if err != nil {
		slog.Error("failed to prepare query to update product", err, slog.Any("product", p))
		return e.ErrorQueryToUpdateProductIsInvalid
	}

	_, err = query.Exec(p.Name, p.Price, p.Description, p.AvailableQuantity, p.Id)
	if err != nil {
		slog.Error("failed to update product", err, slog.Any("product", p))
		return e.ErrorUnableToUpdateProduct
	}

	return nil
}
