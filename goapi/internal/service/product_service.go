package service

import (
	"github.com/gabselbach/imersao17/goapi/internal/database"
	"github.com/gabselbach/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(ProductDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: ProductDB}
}

func (cs *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (cs *ProductService) CreateProduct(name string, description string, price float64, categoryId string, imageUrl string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryId, imageUrl)
	_, err := cs.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (cs *ProductService) FindProduct(id string) (*entity.Product, error) {
	product, err := cs.ProductDB.FindProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (cs *ProductService) GetProductByCategoryID(categoryId string) ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProductByCategoryID(categoryId)
	if err != nil {
		return nil, err
	}
	return products, nil
}
