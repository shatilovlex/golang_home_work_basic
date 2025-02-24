package shopendpoint

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/infrastructure/server/handler/shopendpoint/helper"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
)

type UserEndpoint struct {
	useCase usecase.ShopUsersUseCaseInterface
}

func NewUserEndpoint(useCase usecase.ShopUsersUseCaseInterface) *UserEndpoint {
	return &UserEndpoint{useCase: useCase}
}

func (e UserEndpoint) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var res []*entity.User

	params, err := helper.GetLimitParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error param: %v", err)
	}

	res, err = e.useCase.GetUsers(params)
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		log.Printf("Error getting users: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error encoding user response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (e UserEndpoint) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		userCreateParams repository.UserCreateParams
		users            *entity.User
	)

	err := json.NewDecoder(r.Body).Decode(&userCreateParams)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users, err = e.useCase.CreateUser(userCreateParams)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Printf("Error create user: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println("Error encoding user response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (e UserEndpoint) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		userUpdateParams repository.UserUpdateParams
		res              *entity.User
		err              error
	)

	err = json.NewDecoder(r.Body).Decode(&userUpdateParams)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err = e.useCase.UpdateUser(userUpdateParams)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		log.Printf("Error update user: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error encoding user response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (e UserEndpoint) MakeHandler(r *http.ServeMux) {
	r.Handle("/users", e.handle())
}

func (e UserEndpoint) handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			e.CreateUserHandler(w, r)
		case http.MethodGet:
			e.GetUsersHandler(w, r)
		case http.MethodPut:
			e.UpdateUserHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
