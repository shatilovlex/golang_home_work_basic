package repository

import (
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
)

type ShopProductRepositoryInterface interface {
	Products(arg entity.Params) ([]*entity.Product, error)
	CreateProduct(arg entity.ProductCreateParams) (*entity.Product, error)
	GetProductByID(id int32) (*entity.Product, error)
}
