package usecase

import (
	"testing"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type ShopUserRepositoryStub struct {
	shopUsers map[int32]*entity.ShopUser
	index     int32
}

func NewShopUserRepositoryStub(shopUsers []*entity.ShopUser) ShopUserRepositoryStub {
	var idx int32
	resMap := make(map[int32]*entity.ShopUser, len(shopUsers))
	for _, shopUser := range shopUsers {
		resMap[shopUser.ID] = shopUser
		if shopUser.ID > idx {
			idx = shopUser.ID
		}
	}
	return ShopUserRepositoryStub{
		shopUsers: resMap,
		index:     idx,
	}
}

func (s ShopUserRepositoryStub) Users(_ entity.Params) ([]*entity.ShopUser, error) {
	shopUsers := []*entity.ShopUser{}

	for _, shopUser := range s.shopUsers {
		shopUsers = append(shopUsers, shopUser)
	}

	return shopUsers, nil
}

func (s ShopUserRepositoryStub) UserCreate(arg entity.UserCreateParams) (*entity.ShopUser, error) {
	s.index++
	shopUser := &entity.ShopUser{
		ID:       s.index,
		Name:     arg.Name,
		Email:    arg.Email,
		Password: arg.Password,
	}
	s.shopUsers[s.index] = shopUser
	return shopUser, nil
}

func (s ShopUserRepositoryStub) UserUpdate(arg entity.UserUpdateParams) (*entity.ShopUser, error) {
	s.shopUsers[arg.ID] = &entity.ShopUser{
		Name: arg.Name,
	}

	return s.shopUsers[arg.ID], nil
}

func TestShopUsersUseCase_GetUsers(t *testing.T) {
	e := &entity.ShopUser{
		ID:       1,
		Name:     "Name",
		Email:    "email@mail.loc",
		Password: "password",
	}
	shopUsers := make([]*entity.ShopUser, 0)
	shopUsers = append(shopUsers, e)
	repo := NewShopUserRepositoryStub(shopUsers)
	uc := ShopUsersUseCase{
		repo: repo,
	}
	arg := entity.Params{
		Limit:  10,
		Offset: 0,
	}
	got, err := uc.GetUsers(arg)

	assert.NoError(t, err)
	assert.Equal(t, shopUsers, got)
}

func TestShopUsersUseCase_CreateUser(t *testing.T) {
	shopUsers := make([]*entity.ShopUser, 0)
	repo := NewShopUserRepositoryStub(shopUsers)
	uc := ShopUsersUseCase{
		repo: repo,
	}
	arg := entity.UserCreateParams{
		Name:     "name",
		Email:    "email@mail.loc",
		Password: "password",
	}
	got, err := uc.CreateUser(arg)

	assert.NoError(t, err)
	assert.Equal(t, "name", got.Name)
	assert.Equal(t, "email@mail.loc", got.Email)
	assert.Equal(t, "password", got.Password)
}

func TestShopUsersUseCase_UpdateUser(t *testing.T) {
	e := &entity.ShopUser{
		ID:       1,
		Name:     "Name",
		Email:    "email@mail.loc",
		Password: "password",
	}
	shopUsers := make([]*entity.ShopUser, 0)
	shopUsers = append(shopUsers, e)
	repo := NewShopUserRepositoryStub(shopUsers)
	uc := ShopUsersUseCase{
		repo: repo,
	}
	arg := entity.UserUpdateParams{
		ID:   1,
		Name: "New Name",
	}
	got, err := uc.UpdateUser(arg)

	assert.NoError(t, err)
	assert.Equal(t, "New Name", got.Name)
}
