package usecase

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/db"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
)

type ShopUsersUseCaseInterface interface {
	GetUsers(arg repository.Params) ([]*db.UsersRow, error)
	CreateUser(arg repository.UserCreateParams) (int32, error)
}

type ShopUsersUseCase struct {
	repo repository.ShopUserRepositoryInterface
}

func NewShopUsersUseCase(repo repository.ShopUserRepositoryInterface) *ShopUsersUseCase {
	return &ShopUsersUseCase{repo: repo}
}

func (uc ShopUsersUseCase) GetUsers(arg repository.Params) ([]*db.UsersRow, error) {
	p := repository.Params{
		Limit:  10,
		Offset: 0,
	}
	if arg.Limit > 0 {
		p.Limit = arg.Limit
	} else {
		return nil, fmt.Errorf("limit must be greater than zero")
	}
	if arg.Offset >= 0 {
		p.Offset = arg.Offset
	} else {
		return nil, fmt.Errorf("offset must be greater or equal to zero")
	}

	return uc.repo.Users(p)
}
func (uc ShopUsersUseCase) CreateUser(arg repository.UserCreateParams) (int32, error) {
	return uc.repo.UserCreate(arg)
}
