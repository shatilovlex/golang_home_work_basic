package handler

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_createUser(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/create-user",
		struct{ io.Reader }{strings.NewReader("{\"id\":2,\"name\":\"Jane Doe\",\"age\":18}\n")},
	)
	w := httptest.NewRecorder()

	h := &Handler{}
	h.createUser(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "{\"id\":2,\"name\":\"Jane Doe\",\"age\":18}\n", string(body))
}

func TestHandler_getUser(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/v1/get-user", nil)
	w := httptest.NewRecorder()

	h := &Handler{}
	h.getUser(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, "{\"id\":1,\"name\":\"John Doe\",\"age\":30}\n", string(body))
}

func TestHandler_hello(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/v1/hello", nil)
	w := httptest.NewRecorder()

	h := &Handler{}
	h.hello(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "hello world", string(body))
}
