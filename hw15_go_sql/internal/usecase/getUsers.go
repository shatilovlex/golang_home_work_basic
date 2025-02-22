package usecase

import (
	"context"
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
)

type GetUsersUseCaseInterface interface {
	GetUsers(ctx context.Context, arg Params) ([]*repository.UsersRow, error)
}

type GetUsersUseCase struct {
	querier repository.Querier
}

func NewGetUsersUseCase(querier repository.Querier) *GetUsersUseCase {
	return &GetUsersUseCase{querier: querier}
}

type Params struct {
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}

func (uc GetUsersUseCase) GetUsers(ctx context.Context, arg Params) ([]*repository.UsersRow, error) {
	p := Params{
		Limit:  10,
		Offset: 0,
	}
	if arg.Limit <= 0 {
		return nil, fmt.Errorf("limit must be greater than zero")
	} else {
		p.Limit = arg.Limit
	}
	if arg.Offset < 0 {
		return nil, fmt.Errorf("offset must be greater or equal to zero")
	} else {
		p.Offset = arg.Offset
	}

	params := repository.UsersParams{
		Limit:  p.Limit,
		Offset: p.Offset,
	}

	return uc.querier.Users(ctx, params)
}
