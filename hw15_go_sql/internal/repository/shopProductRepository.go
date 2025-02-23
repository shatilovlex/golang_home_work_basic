package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/db"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
)

type ShopProductRepository struct {
	ctx     context.Context
	querier *db.Queries
	connect *pgxpool.Pool
}

func NewShopProductRepository(ctx context.Context, connect *pgxpool.Pool) ShopProductRepository {
	querier := db.New(connect)
	return ShopProductRepository{ctx: ctx, querier: querier, connect: connect}
}

func (u ShopProductRepository) getProduct(id int32) (*entity.Product, error) {
	item := entity.Product{}
	err := u.connect.QueryRow(
		u.ctx,
		"select id, name, price from pg_storage.shop.users where id = $1 limit 1",
		id,
	).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (u ShopProductRepository) Products(arg entity.Params) ([]*entity.Product, error) {
	var (
		rows pgx.Rows
		err  error
	)
	rows, err = u.connect.Query(u.ctx, db.Products, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*entity.Product{}
	for rows.Next() {
		var i entity.Product
		if err = rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
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
