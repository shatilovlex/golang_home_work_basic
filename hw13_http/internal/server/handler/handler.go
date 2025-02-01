package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

type User struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func (h *Handler) hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	user := User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	fmt.Printf("New user: %+v\n", newUser)

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func (h Handler) InitMux() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/v1/hello", http.HandlerFunc(h.hello))
	mux.Handle("/v1/get-user", http.HandlerFunc(h.getUser))
	mux.Handle("/v1/create-user", http.HandlerFunc(h.createUser))

	return mux
}
