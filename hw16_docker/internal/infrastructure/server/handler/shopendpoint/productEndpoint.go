package shopendpoint

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shatilovlex/golang_home_work_basic/hw16_docker/internal/domain/shop/entity"
	"github.com/shatilovlex/golang_home_work_basic/hw16_docker/internal/domain/shop/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw16_docker/internal/infrastructure/server/handler/shopendpoint/helper"
	"github.com/shatilovlex/golang_home_work_basic/hw16_docker/internal/usecase"
)

type ProductEndpoint struct {
	useCase usecase.ShopProductUseCaseInterface
}

func NewProductEndpoint(useCase usecase.ShopProductUseCaseInterface) *ProductEndpoint {
	return &ProductEndpoint{useCase: useCase}
}

func (e *ProductEndpoint) GetProductHandler(w http.ResponseWriter, r *http.Request) {
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

func (e *ProductEndpoint) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var (
		productCreateParams repository.ProductCreateParams
		products            *entity.Product
	)

	err := json.NewDecoder(r.Body).Decode(&productCreateParams)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	products, err = e.useCase.CreateProduct(productCreateParams)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		log.Printf("Error create product: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (e *ProductEndpoint) MakeHandler(r *http.ServeMux) {
	r.Handle("/products", e.handle())
}

func (e *ProductEndpoint) handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			e.CreateProductHandler(w, r)
		case http.MethodGet:
			e.GetProductHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
