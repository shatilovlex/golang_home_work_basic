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

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/config"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/repository"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/server/handler"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/server/handler/shopEndpoint"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/internal/usecase"
	"github.com/shatilovlex/golang_home_work_basic/hw15_go_sql/pkg/pgconnect"
)

type App struct {
	ctx    context.Context
	config *config.Cfg
	db     *pgxpool.Pool
}

func NewApp(ctx context.Context) (*App, error) {
	conf, err := config.Init()
	if err != nil {
		return nil, err
	}

	db, err := pgconnect.NewDB(ctx, conf.DB)
	if err != nil {
		return nil, err
	}

	return &App{
		ctx:    ctx,
		config: conf,
		db:     db,
	}, nil
}

func (a *App) Start() {
	ctx, stop := signal.NotifyContext(a.ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	querier := repository.New(a.db)

	ip := flag.String("ip", a.config.HTTP.Host, "IP address")
	port := flag.String("port", a.config.HTTP.Port, "Port number")
	flag.Parse()

	ucGetUsers := usecase.NewGetUsersUseCase(querier)
	endpoint := shopEndpoint.NewGetUsersEndpoint(ctx, ucGetUsers)

	addr := fmt.Sprintf("%v:%v", *ip, *port)
	server := &http.Server{
		Addr:              addr,
		Handler:           a.InitMux(endpoint),
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

func (a *App) InitMux(endpoint shopEndpoint.ShopEndpoint) http.Handler {
	mux := http.NewServeMux()

	handler.MakeShopHandlers(mux, endpoint)

	return mux
}
