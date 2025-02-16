package handler

import (
	"context"
	"encoding/json"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
	"net/http"
	"strconv"
)

type Handler struct {
	ctx  context.Context
	repo repository.Querier
}

func New(c context.Context, r repository.Querier) *Handler {
	return &Handler{
		ctx:  c,
		repo: r,
	}
}

func (h Handler) InitMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		h.Users(w, r)
	})

	return mux
}

func (h Handler) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var (
		limit  int64 = 10
		offset int64 = 0
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

	params := usecase.Params{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.New(h.repo)
	res, err := uc.Users(h.ctx, params)

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
