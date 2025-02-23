package repository

import (
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
)

type ShopUserRepositoryInterface interface {
	Users(arg entity.Params) ([]*entity.User, error)
	UserCreate(arg entity.UserCreateParams) (*entity.User, error)
	UserUpdate(arg entity.UserUpdateParams) (*entity.User, error)
}
