package handler

import (
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h Handler) InitMux() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(h.method))

	return mux
}

func (h Handler) method(_ http.ResponseWriter, _ *http.Request) {}
