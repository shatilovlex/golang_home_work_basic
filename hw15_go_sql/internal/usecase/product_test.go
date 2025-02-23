package usecase

import (
	"testing"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/stretchr/testify/assert"
)

type ProductProductRepositoryStub struct {
	products map[int32]*entity.Product
	index    int32
}

func NewProductRepositoryStub(products []*entity.Product) ProductProductRepositoryStub {
	var idx int32
	resMap := make(map[int32]*entity.Product, len(products))
	for _, product := range products {
		resMap[product.ID] = product
		if product.ID > idx {
			idx = product.ID
		}
	}
	return ProductProductRepositoryStub{
		products: resMap,
		index:    idx,
	}
}

func (s ProductProductRepositoryStub) Products(_ entity.Params) ([]*entity.Product, error) {
	products := []*entity.Product{}

	for _, product := range s.products {
		products = append(products, product)
	}

	return products, nil
}

func (s ProductProductRepositoryStub) CreateProduct(arg entity.ProductCreateParams) (*entity.Product, error) {
	s.index++
	product := &entity.Product{
		ID:    s.index,
		Name:  arg.Name,
		Price: arg.Price,
	}
	s.products[s.index] = product
	return product, nil
}

func TestShopUsersUseCase_GetProducts(t *testing.T) {
	e := &entity.Product{
		ID:    1,
		Name:  "Name",
		Price: 10,
	}
	products := make([]*entity.Product, 0)
	products = append(products, e)
	repo := NewProductRepositoryStub(products)
	uc := ShopProductUseCase{
		repo: repo,
	}
	arg := entity.Params{
		Limit:  10,
		Offset: 0,
	}
	got, err := uc.GetProducts(arg)

	assert.NoError(t, err)
	assert.Equal(t, products, got)
}

func TestShopUsersUseCase_CreateProduct(t *testing.T) {
	products := make([]*entity.Product, 0)
	repo := NewProductRepositoryStub(products)
	uc := ShopProductUseCase{
		repo: repo,
	}
	arg := entity.ProductCreateParams{
		Name:  "name",
		Price: 10.0,
	}
	got, err := uc.CreateProduct(arg)

	assert.NoError(t, err)
	assert.Equal(t, "name", got.Name)
	assert.Equal(t, 10.0, got.Price)
}
