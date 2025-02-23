package entity

import "errors"

type Product struct {
	ID    int32   `db:"id" json:"id"`
	Name  string  `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
}

type ProductCreateParams struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var ErrProductNotFound = errors.New("product not found")
