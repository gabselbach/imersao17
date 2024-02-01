package database

import (
	"database/sql"

	"github.com/gabselbach/imersao17/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (cd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := cd.db.Query("Select id, name, description, price, category_id, image_url from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close() // vai rodar só no final

	var producties []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		producties = append(producties, &product)
	}
	return producties, nil
}

func (cd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := cd.db.Query("SELECT id, name,description, price, category_id, image_url FROM products WHERE category_id = ?", categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close() // vai rodar só no final

	var producties []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		producties = append(producties, &product)
	}
	return producties, nil
}

func (cd *ProductDB) FindProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := cd.db.QueryRow("SELECT id, name,description, price, category_id, image_url FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (cd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := cd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?,?,?,?,?,? )",
		product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return nil, err
	}
	return product, nil
}
