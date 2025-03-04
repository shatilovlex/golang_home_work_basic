package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/golang_home_work_basic/hw16_docker/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw16_docker/internal/domain/shop/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw16_docker/internal/infrastructure/db"
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

func (r ShopUserRepository) GetUserByID(id int32) (*entity.User, error) {
	item := entity.User{}
	err := r.connect.QueryRow(
		r.ctx,
		"select id, name, email, password from pg_storage.shop.users where id = $1 limit 1",
		id,
	).Scan(&item.ID, &item.Name, &item.Email, &item.Password)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r ShopUserRepository) UserCreate(arg repository.UserCreateParams) (*entity.User, error) {
	var (
		id  int32
		err error
	)

	params := db.UserCreateParams{
		Name:     &arg.Name,
		Email:    &arg.Email,
		Password: &arg.Password,
	}

	id, err = r.querier.UserCreate(r.ctx, params)
	if err != nil {
		return nil, err
	}

	return r.GetUserByID(id)
}

func (r ShopUserRepository) Users(arg repository.Params) ([]*entity.User, error) {
	var (
		rows pgx.Rows
		err  error
	)
	rows, err = r.connect.Query(r.ctx, db.Users, arg.Limit, arg.Offset)
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
		i.Password = "secret"
		items = append(items, &i)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r ShopUserRepository) UserUpdate(arg repository.UserUpdateParams) (*entity.User, error) {
	var (
		id  int32
		err error
	)

	params := db.UpdateUserNameParams{
		ID:   arg.ID,
		Name: &arg.Name,
	}

	id, err = r.querier.UpdateUserName(r.ctx, params)
	if err != nil {
		return nil, err
	}

	return r.GetUserByID(id)
}
