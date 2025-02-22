package shopEndpoint

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
)

type ShopEndpoint interface {
	GetUsers(ctx context.Context, arg usecase.Params) ([]*repository.UsersRow, error)
	GetUsersHandler(w http.ResponseWriter, r *http.Request)
}

type getUsersEndpoint struct {
	ctx     context.Context
	useCase usecase.GetUsersUseCaseInterface
}

func NewGetUsersEndpoint(ctx context.Context, useCase usecase.GetUsersUseCaseInterface) ShopEndpoint {
	return &getUsersEndpoint{
		ctx:     ctx,
		useCase: useCase,
	}
}

func (e *getUsersEndpoint) GetUsers(ctx context.Context, arg usecase.Params) ([]*repository.UsersRow, error) {
	usersRow, err := e.useCase.GetUsers(ctx, arg)

	if err != nil {
		return nil, err
	}

	return usersRow, nil
}

func (e *getUsersEndpoint) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var (
		limit  int64 = 10
		offset int64 = 0
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

	res, err := e.GetUsers(e.ctx, params)

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
