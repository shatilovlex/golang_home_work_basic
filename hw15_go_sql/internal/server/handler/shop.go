package handler

import (
	"net/http"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/server/handler/shopendpoint"
)

func MakeShopHandlers(r *http.ServeMux, service shopendpoint.Shopendpoint) {
	r.Handle("/users", shopUsers(service))
}

func shopUsers(service shopendpoint.Shopendpoint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			service.CreateUserHandler(w, r)
		case "GET":
			service.GetUsersHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
