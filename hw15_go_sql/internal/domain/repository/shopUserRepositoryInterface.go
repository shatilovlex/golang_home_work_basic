package repository

import (
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
)

type ShopUserRepositoryInterface interface {
	Users(arg entity.Params) ([]*entity.ShopUser, error)
	UserCreate(arg entity.UserCreateParams) (*entity.ShopUser, error)
	UserUpdate(arg entity.UserUpdateParams) (*entity.ShopUser, error)
}
