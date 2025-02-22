package handler

import (
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/server/handler/shopEndpoint"
	"net/http"
)

func MakeShopHandlers(r *http.ServeMux, service shopEndpoint.ShopEndpoint) {
	r.Handle("/users", getShopUsers(service))
}

func getShopUsers(service shopEndpoint.ShopEndpoint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service.GetUsersHandler(w, r)
	})
}
