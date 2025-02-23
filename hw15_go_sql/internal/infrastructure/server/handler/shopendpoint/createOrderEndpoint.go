package shopendpoint

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
)

type CreateOrderEndpoint interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type getCreateOrderEndpoint struct {
	useCase usecase.OrderUseCaseInterface
}

func NewCreateOrderEndpoint(useCase usecase.OrderUseCaseInterface) CreateOrderEndpoint {
	return &getCreateOrderEndpoint{
		useCase: useCase,
	}
}

func (e *getCreateOrderEndpoint) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var createOrderParams entity.CreateOrderParams

	err := json.NewDecoder(r.Body).Decode(&createOrderParams)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = e.useCase.CreateOrder(createOrderParams)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		log.Printf("Error create order: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Order created"})
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
}
