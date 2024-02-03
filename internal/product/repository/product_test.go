package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JPauloMoura/controle-de-estoque/internal/product/entity"
	e "github.com/JPauloMoura/controle-de-estoque/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func createMockDb() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("failed to create mock db:", err)
	}

	return db, mock
}

func createProductTableTitleRow() *sqlmock.Rows {
	collumsProductTable := []string{"id", "name", "price", "description", "availableQuantity"}
	return sqlmock.NewRows(collumsProductTable)
}

func addProductInNewRow(rows *sqlmock.Rows, p entity.Product) *sqlmock.Rows {
	return rows.AddRow(
		p.Id,
		p.Name,
		p.Price,
		p.Description,
		p.AvailableQuantity,
	)
}

func Test_repository_GetProduct(t *testing.T) {
	t.Run("should return error when the query us invalid", func(t *testing.T) {
		productExisted := 2
		invalidQuery := "SELECT \\* FROM  WHERE id=\\$1"

		db, mock := createMockDb()
		defer db.Close()

		mock.
			ExpectQuery(invalidQuery).
			WithArgs(productExisted)

		repo := NewProductRepository(db)
		product, err := repo.GetProduct(productExisted)
		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("should return ErrorProductNotFound when the product id does not exist", func(t *testing.T) {
		productNoExisted := 2

		db, mock := createMockDb()
		defer db.Close()

		rowsEmpty := createProductTableTitleRow()
		mock.
			ExpectQuery("SELECT \\* FROM products WHERE id=\\$1").
			WithArgs(productNoExisted).
			WillReturnRows(rowsEmpty)

		repo := NewProductRepository(db)
		product, err := repo.GetProduct(productNoExisted)

		assert.Equal(t, e.ErrorProductNotFound, err)
		assert.Nil(t, product)
	})

	t.Run("should return error during Scan function when the database item is different from the product entity", func(t *testing.T) {
		productExisted := 1

		db, mock := createMockDb()
		defer db.Close()

		rows := mock.NewRows([]string{"uid", "productName"})
		rows.AddRow(123, "Abancate")

		mock.
			ExpectQuery("SELECT \\* FROM products WHERE id=\\$1").
			WithArgs(productExisted).
			WillReturnRows(rows)

		repo := NewProductRepository(db)
		product, err := repo.GetProduct(productExisted)

		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("should return product successfully", func(t *testing.T) {
		productExisted := 1
		expected := entity.Product{
			Id:                productExisted,
			Name:              "batata",
			Price:             float64(3.5),
			Description:       "detalhes",
			AvailableQuantity: 10,
		}

		db, mock := createMockDb()
		defer db.Close()

		rows := addProductInNewRow(createProductTableTitleRow(), expected)
		mock.
			ExpectQuery("SELECT \\* FROM products WHERE id=\\$1").
			WithArgs(productExisted).
			WillReturnRows(rows)

		repo := NewProductRepository(db)
		product, err := repo.GetProduct(productExisted)

		assert.Nil(t, err)
		assert.Equal(t, &expected, product)
	})

}
