package repository

import (
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
)

type ShopProductRepositoryInterface interface {
	Products(arg Params) ([]*entity.Product, error)
	CreateProduct(arg ProductCreateParams) (*entity.Product, error)
	GetProductByID(id int32) (*entity.Product, error)
}

type ProductCreateParams struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
