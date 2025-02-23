package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/db"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
)

type Repository struct {
	ctx     context.Context
	querier *db.Queries
	connect *pgxpool.Pool
}

func NewRepository(ctx context.Context, connect *pgxpool.Pool) Repository {
	querier := db.New(connect)
	return Repository{ctx: ctx, querier: querier, connect: connect}
}

func (u Repository) UserCreate(arg entity.UserCreateParams) (int32, error) {
	params := db.UserCreateParams{
		Name:     &arg.Name,
		Email:    &arg.Email,
		Password: &arg.Password,
	}

	return u.querier.UserCreate(u.ctx, params)
}

func (u Repository) Users(arg entity.Params) ([]*entity.ShopUser, error) {
	var (
		rows pgx.Rows
		err  error
	)
	rows, err = u.connect.Query(u.ctx, db.Users, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*entity.ShopUser{}
	for rows.Next() {
		var i entity.ShopUser
		if err = rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
