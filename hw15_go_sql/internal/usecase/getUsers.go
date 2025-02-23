package usecase

import (
	"context"
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
)

type ShopUsersUseCaseInterface interface {
	GetUsers(ctx context.Context, arg Params) ([]*repository.UsersRow, error)
	CreateUser(ctx context.Context, arg UserCreateParams) (int32, error)
}

type ShopUsersUseCase struct {
	querier repository.Querier
}

func NewShopUsersUseCase(querier repository.Querier) *ShopUsersUseCase {
	return &ShopUsersUseCase{querier: querier}
}

type Params struct {
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}

func (uc ShopUsersUseCase) GetUsers(ctx context.Context, arg Params) ([]*repository.UsersRow, error) {
	p := Params{
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

	params := repository.UsersParams{
		Limit:  p.Limit,
		Offset: p.Offset,
	}

	return uc.querier.Users(ctx, params)
}

type UserCreateParams struct {
	Name     *string `db:"name" json:"name"`
	Email    *string `db:"email" json:"email"`
	Password *string `db:"password" json:"password"`
}

func (uc ShopUsersUseCase) CreateUser(ctx context.Context, arg UserCreateParams) (int32, error) {
	params := repository.UserCreateParams{
		Name:     arg.Name,
		Email:    arg.Email,
		Password: arg.Password,
	}
	return uc.querier.UserCreate(ctx, params)
}
