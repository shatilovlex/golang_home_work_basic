package usecase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/infrastructure/db"
)

type OrderUseCaseInterface interface {
	CreateOrder(params entity.CreateOrderParams) error
}

type ShopOrderUseCase struct {
	ctx               context.Context
	connect           *pgxpool.Pool
	userRepository    repository.ShopUserRepositoryInterface
	productRepository repository.ShopProductRepositoryInterface
}

func NewShopOrderUseCase(
	ctx context.Context,
	connect *pgxpool.Pool,
	userRepository repository.ShopUserRepositoryInterface,
	productRepository repository.ShopProductRepositoryInterface,
) *ShopOrderUseCase {
	return &ShopOrderUseCase{
		ctx:               ctx,
		connect:           connect,
		userRepository:    userRepository,
		productRepository: productRepository,
	}
}

func (uc *ShopOrderUseCase) CreateOrder(params entity.CreateOrderParams) error {
	tx, err := uc.connect.BeginTx(uc.ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback(uc.ctx)

	_, err = uc.userRepository.GetUserByID(params.UserID)
	if err != nil {
		return entity.ErrUserNotFound
	}
	var (
		totalAmount float64
		orderID     int32
	)

	totalAmount, err = uc.getTotalAmount(params.ProductIds)
	if err != nil {
		return err
	}

	orderID, err = uc.NewOrder(params.UserID, totalAmount)
	if err != nil {
		return err
	}

	err = uc.MassInsertOrderProducts(orderID, params.ProductIds)
	if err != nil {
		return err
	}

	err = tx.Commit(uc.ctx)

	if err != nil {
		return err
	}

	return nil
}

func (uc *ShopOrderUseCase) getTotalAmount(productIds []int32) (float64, error) {
	var totalAmount float64

	for _, productID := range productIds {
		product, err := uc.productRepository.GetProductByID(productID)
		if err != nil {
			return 0, entity.ErrProductNotFound
		}

		totalAmount += product.Price
	}
	return totalAmount, nil
}

func (uc *ShopOrderUseCase) NewOrder(userID int32, amount float64) (int32, error) {
	row := uc.connect.QueryRow(
		uc.ctx,
		db.CreateOrder,
		userID,
		pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		amount,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

func (uc *ShopOrderUseCase) MassInsertOrderProducts(orderID int32, productIds []int32) error {
	for _, productID := range productIds {
		_, err := uc.connect.Exec(uc.ctx, db.CreateOrderProduct, orderID, productID)
		if err != nil {
			return err
		}
	}
	return nil
}
