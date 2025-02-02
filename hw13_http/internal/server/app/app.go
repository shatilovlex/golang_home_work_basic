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

	"github.com/shatilovlex/golang_home_work_basic/hw13_http/internal/server/handler"
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

	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.String("port", "8080", "Port number")
	flag.Parse()

	h := handler.New()
	addr := fmt.Sprintf("%v:%v", *ip, *port)
	server := &http.Server{
		Addr:              addr,
		Handler:           middlewareLogToConsole(h.InitMux()),
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

func middlewareLogToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
