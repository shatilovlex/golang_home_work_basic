package app

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type API struct {
	Client *http.Client
	Method string
	URL    string
	Body   string
}

func (api *API) DoStuff() ([]byte, error) {
	request, err := http.NewRequestWithContext(context.Background(), api.Method, api.URL, strings.NewReader(api.Body))
	if err != nil {
		return nil, fmt.Errorf("ошибка формирования запроса: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")

	resp, err := api.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка HTTP-ответа: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
