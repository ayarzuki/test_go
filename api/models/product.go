package models

import (
	"beauty_store/api/database"
)

type Product struct {
	ID_PRODUCT   int     `json:"id_product"`
	NAME_PRODUCT string  `json:"name_product"`
	PRICE        float64 `json:"price"`
}

func GetAllProducts() ([]Product, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID_PRODUCT, &product.NAME_PRODUCT, &product.PRICE)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
