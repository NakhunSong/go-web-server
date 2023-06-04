package services

import (
	"nakhun/web-server/models"
)

func GetProducts() []models.Product {
	products, err := models.GetProducts()
	if err != nil {
		return []models.Product{}
	}
	return products
}

func CreateProduct(title string, price int64) (int64, error) {
	return models.CreateProduct(title, price)
}
