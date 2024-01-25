package models

import (
	"log"
	"strconv"

	"github.com/14-web_api/data"
)

// Products is the representation of a product entity available for sale
type Product struct {
	Id                int
	Name              string
	Price             float64
	Description       string
	AvailableQuantity int
}

func GetAllProducts() []Product {
	db := data.ConnectDb()
	defer db.Close()

	query, err := db.Query("SELECT * FROM products ORDER BY name ASC")
	if err != nil {
		log.Println("failed to get products: ", err)
		return []Product{}
	}

	listProduct := []Product{}

	for query.Next() {
		var id, availableQuantity int
		var name, description string
		var price float64

		// its parameters must be in the same order as the entity fields
		err := query.Scan(&id, &name, &price, &description, &availableQuantity)
		if err != nil {
			log.Println("failed to scan products: ", err)
			return []Product{}
		}

		p := Product{id, name, float64(price), description, int(availableQuantity)}

		listProduct = append(listProduct, p)
	}
	return listProduct
}

func InsertProduct(p Product) {
	db := data.ConnectDb()
	defer db.Close()

	queryInsert, err := db.Prepare(`
		insert into products(name, description, price, available_quantity)
		values($1, $2, $3, $4)
	`)

	if err != nil {
		panic(err.Error())
	}

	_, err = queryInsert.Exec(p.Name, p.Description, p.Price, p.AvailableQuantity)
	if err != nil {
		log.Println("failed to insert product: ", err)
	}

}

func DeleteProduct(id string) {
	db := data.ConnectDb()
	defer db.Close()

	query, err := db.Prepare(`
		DELETE FROM products WHERE id=$1
	`)
	if err != nil {
		log.Println("failed to prepare query to delete product: ", err)
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to delete product, id is not valid: ", err)
		return
	}

	_, err = query.Exec(uid)

	if err != nil {
		log.Println("failed to delete product: ", err)
		return
	}

	log.Println("product deleted! id ", uid)
}

func GetProduct(id string) Product {
	db := data.ConnectDb()
	defer db.Close()

	intID, err := strconv.Atoi(string(id))
	if err != nil {
		log.Println("failed to convert string id to integer: " + err.Error())
		return Product{}
	}

	query, err := db.Query(`SELECT * FROM products WHERE id=$1`, intID)
	if err != nil {
		log.Println("failed to get product by id: " + err.Error())
		return Product{}
	}

	p := Product{}

	for query.Next() {
		// its parameters must be in the same order as the entity fields
		err := query.Scan(&p.Id, &p.Name, &p.Price, &p.Description, &p.AvailableQuantity)
		if err != nil {
			log.Println("failed to scan when get product by id: ", err)
			return Product{}
		}
	}

	return p
}

func UpdateProduct(p Product) {
	db := data.ConnectDb()
	defer db.Close()

	// product := GetProduct(string(p.Id))

	query, err := db.Prepare(`
		UPDATE products SET 
		name=$1, price=$2, description=$3, available_quantity=$4
		WHERE id=$5
	`)
	if err != nil {
		log.Println("failed to prepare query to update product: ", err)
		return
	}

	_, err = query.Exec(p.Name, p.Price, p.Description, p.AvailableQuantity, p.Id)

	if err != nil {
		log.Println("failed to update product: ", err)
		return
	}

}
