package handler

import (
	"net/http"
)

type Handler interface {
	MakeHandler(r *http.ServeMux)
}

type Maker struct {
	handlers []Handler
}

func NewMaker(handlers []Handler) *Maker {
	return &Maker{handlers: handlers}
}

func (handlers *Maker) MakeHandlers(mux *http.ServeMux) {
	for _, handler := range handlers.handlers {
		handler.MakeHandler(mux)
	}
}
