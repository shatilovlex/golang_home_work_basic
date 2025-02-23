package shopendpoint

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
)

type UserEndpoint interface {
	GetUsersHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
}

type getUserEndpoint struct {
	useCase usecase.ShopUsersUseCaseInterface
}

func NewUserEndpoint(useCase usecase.ShopUsersUseCaseInterface) UserEndpoint {
	return &getUserEndpoint{
		useCase: useCase,
	}
}

func (e *getUserEndpoint) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("get error: %v", err)
			return
		}
	}
	if offsetRaw != "" {
		offset, err = strconv.ParseInt(offsetRaw, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("get error: %v", err)
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
		log.Printf("get error: %v", err)
		return
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
	}
}

func (e getUserEndpoint) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("get error: %v", err)
		return
	}

	var (
		userCreateParams entity.UserCreateParams
		res              *entity.ShopUser
	)

	err = json.Unmarshal(body, &userCreateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}

	res, err = e.useCase.CreateUser(userCreateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(resBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
	}
}

func (e getUserEndpoint) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("get error: %v", err)
		return
	}

	var (
		updateParams entity.UserUpdateParams
		res          *entity.ShopUser
	)
	err = json.Unmarshal(body, &updateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}

	res, err = e.useCase.UpdateUser(updateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
	}
}
