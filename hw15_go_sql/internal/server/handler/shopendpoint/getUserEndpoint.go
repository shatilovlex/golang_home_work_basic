package shopendpoint

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
)

type Shopendpoint interface {
	GetUsersHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
}

type getShopEndpoint struct {
	useCase usecase.ShopUsersUseCaseInterface
}

func NewShopEndpoint(useCase usecase.ShopUsersUseCaseInterface) Shopendpoint {
	return &getShopEndpoint{
		useCase: useCase,
	}
}

func (e *getShopEndpoint) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var (
		limit  int64 = 10
		offset int64
		res    []*entity.ShopUser
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

	params := entity.Params{
		Limit:  limit,
		Offset: offset,
	}

	res, err = e.useCase.GetUsers(params)
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

	var userCreateParams entity.UserCreateParams

	err = json.Unmarshal(body, &userCreateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = e.useCase.CreateUser(userCreateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
