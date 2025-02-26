package main

import (
	"context"
	"log"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/infrastructure/server/app"
)

func main() {
	ctx := context.Background()
	mainApp, err := app.NewApp(ctx)
	if err != nil {
		log.Panicln(err.Error())
	}
	mainApp.Start()
}
