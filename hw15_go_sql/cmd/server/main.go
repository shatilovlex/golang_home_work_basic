package main

import (
	"context"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/server/app"
)

func main() {
	ctx := context.Background()
	mainApp := app.NewApp(ctx)
	mainApp.Start()
}
