package shopendpoint

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
)

type Shopendpoint interface {
	GetUsers(ctx context.Context, arg usecase.Params) ([]*repository.UsersRow, error)
	GetUsersHandler(w http.ResponseWriter, r *http.Request)
	CreateUser(ctx context.Context, arg usecase.UserCreateParams) (int32, error)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
}

type getShopEndpoint struct {
	ctx     context.Context
	useCase usecase.ShopUsersUseCaseInterface
}

func NewShopEndpoint(ctx context.Context, useCase usecase.ShopUsersUseCaseInterface) Shopendpoint {
	return &getShopEndpoint{
		ctx:     ctx,
		useCase: useCase,
	}
}

func (e *getShopEndpoint) GetUsers(ctx context.Context, arg usecase.Params) ([]*repository.UsersRow, error) {
	return e.useCase.GetUsers(ctx, arg)
}

func (e *getShopEndpoint) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var (
		limit  int64 = 10
		offset int64
		res    []*repository.UsersRow
		err    error
	)
	limitRaw := r.URL.Query().Get("limit")
	offsetRaw := r.URL.Query().Get("offset")

	if limitRaw != "" {
		limit, err = strconv.ParseInt(limitRaw, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if offsetRaw != "" {
		offset, err = strconv.ParseInt(offsetRaw, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	params := usecase.Params{
		Limit:  limit,
		Offset: offset,
	}

	res, err = e.GetUsers(e.ctx, params)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (e *getShopEndpoint) CreateUser(ctx context.Context, arg usecase.UserCreateParams) (int32, error) {
	return e.useCase.CreateUser(ctx, arg)
}

func (e getShopEndpoint) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userCreateParams usecase.UserCreateParams

	err = json.Unmarshal(body, &userCreateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = e.CreateUser(e.ctx, userCreateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
