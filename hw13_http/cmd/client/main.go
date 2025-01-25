package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	method := flag.String("method", "GET", "Request method")
	url := flag.String("url", "http://localhost:8088/v1/get-user", "Url")
	body := flag.String("body", "", "Body")
	flag.Parse()

	client := &http.Client{}
	request, err := http.NewRequestWithContext(context.Background(), *method, *url, strings.NewReader(*body))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	fmt.Println(string(respBody))
}
