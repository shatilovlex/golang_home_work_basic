package usecase

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/repository"
)

type ShopProductUseCaseInterface interface {
	GetProducts(arg entity.Params) ([]*entity.Product, error)
}

type ShopProductUseCase struct {
	repo repository.ShopProductRepositoryInterface
}

func NewShopProductUseCase(repo repository.ShopProductRepositoryInterface) *ShopProductUseCase {
	return &ShopProductUseCase{repo: repo}
}

func (uc ShopProductUseCase) GetProducts(arg entity.Params) ([]*entity.Product, error) {
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

	return uc.repo.Products(p)
}
