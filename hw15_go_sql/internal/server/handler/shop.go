package handler

import (
	"net/http"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/server/handler/shopendpoint"
)

func MakeShopHandlers(r *http.ServeMux, service shopendpoint.UserEndpoint) {
	r.Handle("/users", shopUsers(service))
}

func shopUsers(service shopendpoint.UserEndpoint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			service.CreateUserHandler(w, r)
		case "GET":
			service.GetUsersHandler(w, r)
		case "PUT":
			service.UpdateUserHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func MakeProductHandlers(r *http.ServeMux, service shopendpoint.ProductEndpoint) {
	r.Handle("/products", shopProducts(service))
}

func shopProducts(service shopendpoint.ProductEndpoint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			service.GetProductHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
