package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shatilovlex/golang_home_work_basic/hw13_http/internal/client/app"
	"github.com/stretchr/testify/assert"
)

func TestDoStuffWithTestServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello world"))
	}))
	defer server.Close()

	api := app.API{
		Client: server.Client(),
		Method: "GET",
		URL:    server.URL,
	}
	body, err := api.DoStuff()

	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), body)
}
