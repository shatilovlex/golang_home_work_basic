package shopendpoint

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/domain/shop/entity"
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
		limit  int64 = 10
		offset int64
		res    []*entity.Product
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

	res, err = e.useCase.GetProducts(params)
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

func (e *getProductEndpoint) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("get error: %v", err)
		return
	}

	var (
		userCreateParams entity.ProductCreateParams
		res              *entity.Product
	)

	err = json.Unmarshal(body, &userCreateParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("get error: %v", err)
		return
	}

	res, err = e.useCase.CreateProduct(userCreateParams)
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
