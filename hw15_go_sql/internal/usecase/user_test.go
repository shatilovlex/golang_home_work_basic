package usecase

import (
	"testing"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/repository"
	"github.com/stretchr/testify/assert"
)

type ShopUserRepositoryStub struct {
	shopUsers map[int32]*entity.User
	index     int32
}

func NewShopUserRepositoryStub(shopUsers []*entity.User) ShopUserRepositoryStub {
	var idx int32
	resMap := make(map[int32]*entity.User, len(shopUsers))
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

func (r ShopUserRepositoryStub) GetUserByID(id int32) (*entity.User, error) {
	return r.shopUsers[id], nil
}

func (r ShopUserRepositoryStub) Users(_ repository.Params) ([]*entity.User, error) {
	shopUsers := []*entity.User{}

	for _, shopUser := range r.shopUsers {
		shopUsers = append(shopUsers, shopUser)
	}

	return shopUsers, nil
}

func (r ShopUserRepositoryStub) UserCreate(arg repository.UserCreateParams) (*entity.User, error) {
	r.index++
	shopUser := &entity.User{
		ID:       r.index,
		Name:     arg.Name,
		Email:    arg.Email,
		Password: arg.Password,
	}
	r.shopUsers[r.index] = shopUser
	return shopUser, nil
}

func (r ShopUserRepositoryStub) UserUpdate(arg repository.UserUpdateParams) (*entity.User, error) {
	r.shopUsers[arg.ID] = &entity.User{
		Name: arg.Name,
	}

	return r.shopUsers[arg.ID], nil
}

func TestShopUsersUseCase_GetUsers(t *testing.T) {
	e := &entity.User{
		ID:       1,
		Name:     "Name",
		Email:    "email@mail.loc",
		Password: "password",
	}
	shopUsers := make([]*entity.User, 0)
	shopUsers = append(shopUsers, e)
	repo := NewShopUserRepositoryStub(shopUsers)
	uc := ShopUsersUseCase{
		repo: repo,
	}
	arg := repository.Params{
		Limit:  10,
		Offset: 0,
	}
	got, err := uc.GetUsers(arg)

	assert.NoError(t, err)
	assert.Equal(t, shopUsers, got)
}

func TestShopUsersUseCase_CreateUser(t *testing.T) {
	shopUsers := make([]*entity.User, 0)
	repo := NewShopUserRepositoryStub(shopUsers)
	uc := ShopUsersUseCase{
		repo: repo,
	}
	arg := repository.UserCreateParams{
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
	e := &entity.User{
		ID:       1,
		Name:     "Name",
		Email:    "email@mail.loc",
		Password: "password",
	}
	shopUsers := make([]*entity.User, 0)
	shopUsers = append(shopUsers, e)
	repo := NewShopUserRepositoryStub(shopUsers)
	uc := ShopUsersUseCase{
		repo: repo,
	}
	arg := repository.UserUpdateParams{
		ID:   1,
		Name: "New Name",
	}
	got, err := uc.UpdateUser(arg)

	assert.NoError(t, err)
	assert.Equal(t, "New Name", got.Name)
}
