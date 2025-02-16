package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/config"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/server/handler"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/pkg/pgconnect"
)

type App struct {
	ctx context.Context
}

func NewApp(ctx context.Context) *App {
	return &App{ctx: ctx}
}

func (a *App) Start() {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	conf, err := config.Init()
	if err != nil {
		log.Panicln(err.Error())
	}

	db, err := pgconnect.NewDB(ctx, conf.DB)
	if err != nil {
		log.Panicln(err.Error())
	}
	repo := repository.New(db)

	ip := flag.String("ip", conf.HTTP.Host, "IP address")
	port := flag.String("port", conf.HTTP.Port, "Port number")
	flag.Parse()

	h := handler.New(ctx, repo)
	addr := fmt.Sprintf("%v:%v", *ip, *port)
	server := &http.Server{
		Addr:              addr,
		Handler:           h.InitMux(),
		ReadHeaderTimeout: 2 * time.Second,
	}
	go func() {
		log.Printf("start receiving at: %v", addr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	ctxT, cancel := context.WithTimeout(a.ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxT); err != nil {
		log.Printf("error while shutting down http server: %s", err)
	}
	log.Println("final")
}
