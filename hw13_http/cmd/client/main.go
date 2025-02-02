package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/shatilovlex/golang_home_work_basic/hw13_http/internal/client/app"
)

func main() {
	method := flag.String("method", "GET", "Request method")
	url := flag.String("url", "http://localhost:8080/v1/get-user", "URL")
	body := flag.String("body", "", "Body")
	flag.Parse()

	client := app.API{
		Client: &http.Client{},
		URL:    *url,
		Method: *method,
		Body:   *body,
	}

	respBody, err := client.DoStuff()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(respBody))
}
