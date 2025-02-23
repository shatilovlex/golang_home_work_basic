package repository

import (
	"context"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/db"
)

type ShopUserRepositoryInterface interface {
	Users(arg Params) ([]*db.UsersRow, error)
	UserCreate(arg UserCreateParams) (int32, error)
}

type Repository struct {
	ctx     context.Context
	querier *db.Queries
}

type UserCreateParams struct {
	Name     *string `db:"name" json:"name"`
	Email    *string `db:"email" json:"email"`
	Password *string `db:"password" json:"password"`
}

type Params struct {
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}

func NewRepository(ctx context.Context, querier *db.Queries) Repository {
	return Repository{ctx: ctx, querier: querier}
}

func (u Repository) UserCreate(arg UserCreateParams) (int32, error) {
	params := db.UserCreateParams{
		Name:     arg.Name,
		Email:    arg.Email,
		Password: arg.Password,
	}

	return u.querier.UserCreate(u.ctx, params)
}

func (u Repository) Users(arg Params) ([]*db.UsersRow, error) {
	params := db.UsersParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
	}

	return u.querier.Users(u.ctx, params)
}
