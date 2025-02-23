package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/db"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
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

func (u ShopUserRepository) getUser(id int32) (*entity.ShopUser, error) {
	item := entity.ShopUser{}
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

func (u ShopUserRepository) UserCreate(arg entity.UserCreateParams) (*entity.ShopUser, error) {
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

func (u ShopUserRepository) Users(arg entity.Params) ([]*entity.ShopUser, error) {
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

func (u ShopUserRepository) UserUpdate(arg entity.UserUpdateParams) (*entity.ShopUser, error) {
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
