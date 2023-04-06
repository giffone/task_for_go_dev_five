package app

import (
	"fmt"
	"log"
	"nbrates/internal/api"
	"nbrates/internal/config"
	"nbrates/internal/repository/client"
	"nbrates/internal/service"
	"os"
	"os/signal"
	"syscall"

	"nbrates/internal/repository/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	errCh    chan error
	signalCh chan os.Signal
	conf     *config.AppConf
	router   *fiber.App
	pool     *pgxpool.Pool
}

func New(conf *config.AppConf) *App {
	app := App{
		conf:     conf,
		pool:     newStorage(conf),
		errCh:    make(chan error),
		signalCh: make(chan os.Signal, 1),
	}
	signal.Notify(app.signalCh, syscall.SIGINT, syscall.SIGTERM)

	resty := newCli(conf.Nb.Link)
	cli := client.New(resty)
	db := storage.New(app.pool)
	svc := service.New(db, cli)
	hndl := api.New(svc)
	app.router = newRouter(hndl)
	return &app
}

func (a *App) Start() {
	addr := fmt.Sprintf("%s:%s", a.conf.Route.Host, a.conf.Route.Port)

	go func() {
		a.errCh <- a.router.Listen(addr)
	}()

	select {
	case err := <-a.errCh:
		log.Printf("router listener: %s\n", err.Error())
	case signal := <-a.signalCh:
		log.Printf("signal to stop: %v\n", signal)
	}
}

func (a *App) Stop() {
	log.Println("[graseful stop] stopping...")
	if a.pool != nil {
		a.pool.Close()
	}
	close(a.errCh)
	close(a.signalCh)
}
