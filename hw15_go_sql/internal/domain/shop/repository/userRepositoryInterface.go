package repository

import (
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
)

type ShopUserRepositoryInterface interface {
	Users(arg Params) ([]*entity.User, error)
	UserCreate(arg UserCreateParams) (*entity.User, error)
	UserUpdate(arg UserUpdateParams) (*entity.User, error)
	GetUserByID(id int32) (*entity.User, error)
}

type UserUpdateParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type UserCreateParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
