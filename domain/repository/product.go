package repository

import (
	"database/sql"
	"log/slog"
	"strconv"

	"github.com/JPauloMoura/controle-de-estoque/domain/entity"
)

type ProductRepository interface {
	GetAllProducts() []entity.Product
	InsertProduct(p entity.Product)
	DeleteProduct(id string)
	GetProduct(id string) entity.Product
	UpdateProduct(p entity.Product)
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r repository) GetAllProducts() []entity.Product {
	query, err := r.db.Query("SELECT * FROM products ORDER BY name ASC")
	if err != nil {
		slog.Error("failed to get products: ", err)
		return []entity.Product{}
	}

	listProduct := []entity.Product{}

	for query.Next() {
		var id, availableQuantity int
		var name, description string
		var price float64

		// its parameters must be in the same order as the entity fields
		err := query.Scan(&id, &name, &price, &description, &availableQuantity)
		if err != nil {
			slog.Error("failed to scan products: ", err)
			return []entity.Product{}
		}

		p := entity.Product{Id: id, Name: name, Price: price, Description: description, AvailableQuantity: availableQuantity}

		listProduct = append(listProduct, p)
	}

	return listProduct
}

func (r repository) InsertProduct(p entity.Product) {
	queryInsert, err := r.db.Prepare(`
		insert into products(name, description, price, available_quantity)
		values($1, $2, $3, $4)
	`)

	if err != nil {
		slog.Error("failed to prepare query", err)
		return
	}

	_, err = queryInsert.Exec(p.Name, p.Description, p.Price, p.AvailableQuantity)
	if err != nil {
		slog.Error("failed to insert product", err)
		return
	}

	slog.Debug("product created", slog.String("name", p.Name))
}

func (r repository) DeleteProduct(id string) {
	query, err := r.db.Prepare(`
		DELETE FROM products WHERE id=$1
	`)
	if err != nil {
		slog.Error("failed to prepare query to delete product", err)
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("failed to delete product, id is not valid", err)
		return
	}

	_, err = query.Exec(uid)

	if err != nil {
		slog.Error("failed to delete product", err)
		return
	}

	slog.Debug("product deleted", slog.Int("productId", uid))
}

func (r repository) GetProduct(id string) entity.Product {
	productId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("failed to convert string id to integer", err)
		return entity.Product{}
	}

	query, err := r.db.Query(`SELECT * FROM products WHERE id=$1`, productId)
	if err != nil {
		slog.Error("failed to get product by id", err, slog.Int("productId", productId))
		return entity.Product{}
	}

	p := entity.Product{}

	for query.Next() {
		// its parameters must be in the same order as the entity fields
		err := query.Scan(&p.Id, &p.Name, &p.Price, &p.Description, &p.AvailableQuantity)
		if err != nil {
			slog.Error("failed to scan when get product by id", err, slog.Int("productId", productId))
			return entity.Product{}
		}
	}

	slog.Debug("product found", slog.Int("productId", p.Id))
	return p
}

func (r repository) UpdateProduct(p entity.Product) {
	query, err := r.db.Prepare(`
		UPDATE products SET 
		name=$1, price=$2, description=$3, available_quantity=$4
		WHERE id=$5
	`)
	if err != nil {
		slog.Error("failed to prepare query to update product", err, slog.Any("product", p))
		return
	}

	_, err = query.Exec(p.Name, p.Price, p.Description, p.AvailableQuantity, p.Id)

	if err != nil {
		slog.Error("failed to update product", err, slog.Any("product", p))
		return
	}

	slog.Debug("product updated", slog.Int("productId", p.Id))
}
