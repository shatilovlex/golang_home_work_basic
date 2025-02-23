package repository

import (
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
)

type ShopProductRepositoryInterface interface {
	Products(arg entity.Params) ([]*entity.Product, error)
}
