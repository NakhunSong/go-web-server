package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Product struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Price     int64  `json:"price"`
	CreatedAt string `json:"created_at"`
}

func InitDB() error {
	var err error

	db, err = sql.Open("mysql", "root:12341234@(127.0.0.1:3306)/goweb?parseTime=true")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CloseDB() {
	db.Close()
}

func GetProducts() ([]Product, error) {
	rows, err := db.Query("SELECT * FROM product")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product

		if err := rows.Scan(&product.ID, &product.Title, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func CreateProduct(title string, price int64) (int64, error) {
	result, err := db.Exec("INSERT INTO product (title, price) VALUES (?, ?)", title, price)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
