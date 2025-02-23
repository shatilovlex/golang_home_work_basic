package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/infrastructure/db"
)

type ShopUserRepository struct {
	ctx     context.Context
	querier *db.Queries
	connect *pgxpool.Pool
}

func NewShopUserRepository(ctx context.Context, connect *pgxpool.Pool) ShopUserRepository {
	querier := db.New(connect)
	return ShopUserRepository{ctx: ctx, querier: querier, connect: connect}
}

func (u ShopUserRepository) getUser(id int32) (*entity.User, error) {
	item := entity.User{}
	err := u.connect.QueryRow(
		u.ctx,
		"select id, name, email, password from pg_storage.shop.users where id = $1 limit 1",
		id,
	).Scan(&item.ID, &item.Name, &item.Email, &item.Password)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (u ShopUserRepository) UserCreate(arg entity.UserCreateParams) (*entity.User, error) {
	var (
		id  int32
		err error
	)

	params := db.UserCreateParams{
		Name:     &arg.Name,
		Email:    &arg.Email,
		Password: &arg.Password,
	}

	id, err = u.querier.UserCreate(u.ctx, params)
	if err != nil {
		return nil, err
	}

	return u.getUser(id)
}

func (u ShopUserRepository) Users(arg entity.Params) ([]*entity.User, error) {
	var (
		rows pgx.Rows
		err  error
	)
	rows, err = u.connect.Query(u.ctx, db.Users, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*entity.User{}
	for rows.Next() {
		var i entity.User
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

func (u ShopUserRepository) UserUpdate(arg entity.UserUpdateParams) (*entity.User, error) {
	var (
		id  int32
		err error
	)

	params := db.UpdateUserNameParams{
		ID:   arg.ID,
		Name: &arg.Name,
	}

	id, err = u.querier.UpdateUserName(u.ctx, params)
	if err != nil {
		return nil, err
	}

	return u.getUser(id)
}
