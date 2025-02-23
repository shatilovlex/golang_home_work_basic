package usecase

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/repository"
)

type ShopUsersUseCaseInterface interface {
	GetUsers(arg entity.Params) ([]*entity.User, error)
	CreateUser(arg entity.UserCreateParams) (*entity.User, error)
	UpdateUser(arg entity.UserUpdateParams) (*entity.User, error)
}

type ShopUsersUseCase struct {
	repo repository.ShopUserRepositoryInterface
}

func NewShopUsersUseCase(repo repository.ShopUserRepositoryInterface) *ShopUsersUseCase {
	return &ShopUsersUseCase{repo: repo}
}

func (uc ShopUsersUseCase) GetUsers(arg entity.Params) ([]*entity.User, error) {
	p := entity.Params{
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

func (uc ShopUsersUseCase) CreateUser(arg entity.UserCreateParams) (*entity.User, error) {
	return uc.repo.UserCreate(arg)
}

func (uc ShopUsersUseCase) UpdateUser(arg entity.UserUpdateParams) (*entity.User, error) {
	return uc.repo.UserUpdate(arg)
}
