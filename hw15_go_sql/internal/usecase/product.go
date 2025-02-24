package usecase

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/repository"
)

type ShopProductUseCaseInterface interface {
	GetProducts(arg repository.Params) ([]*entity.Product, error)
	CreateProduct(arg repository.ProductCreateParams) (*entity.Product, error)
}

type ShopProductUseCase struct {
	repo repository.ShopProductRepositoryInterface
}

func NewShopProductUseCase(repo repository.ShopProductRepositoryInterface) *ShopProductUseCase {
	return &ShopProductUseCase{repo: repo}
}

func (uc ShopProductUseCase) GetProducts(arg repository.Params) ([]*entity.Product, error) {
	p := repository.Params{
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

func (uc ShopProductUseCase) CreateProduct(arg repository.ProductCreateParams) (*entity.Product, error) {
	return uc.repo.CreateProduct(arg)
}
