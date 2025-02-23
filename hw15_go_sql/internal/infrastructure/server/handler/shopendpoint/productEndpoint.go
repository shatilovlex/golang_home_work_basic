package shopendpoint

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/infrastructure/server/handler/shopendpoint/helper"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
)

type ProductEndpoint interface {
	GetProductHandler(w http.ResponseWriter, r *http.Request)
	CreateProductHandler(w http.ResponseWriter, r *http.Request)
}

type getProductEndpoint struct {
	useCase usecase.ShopProductUseCaseInterface
}

func NewProductEndpoint(useCase usecase.ShopProductUseCaseInterface) ProductEndpoint {
	return &getProductEndpoint{
		useCase: useCase,
	}
}

func (e *getProductEndpoint) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var (
		res []*entity.Product
		err error
	)
	params, err := helper.GetLimitParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error param: %v", err)
	}

	res, err = e.useCase.GetProducts(params)
	if err != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		log.Printf("Error getting products: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (e *getProductEndpoint) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var (
		productCreateParams entity.ProductCreateParams
		res                 *entity.Product
	)

	err := json.NewDecoder(r.Body).Decode(&productCreateParams)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err = e.useCase.CreateProduct(productCreateParams)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		log.Printf("Error create product: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
