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

	"nbrates/internal/repository/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	quit   chan os.Signal
	conf   *config.AppConf
	router *fiber.App
	pool   *pgxpool.Pool
}

func New(conf *config.AppConf) *App {
	app := App{
		conf: conf,
		pool: newStorage(conf),
		quit: make(chan os.Signal, 1),
	}
	signal.Notify(app.quit, os.Interrupt)

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
		if err := a.router.Listen(addr); err != nil {
			log.Printf("router listener: %s\n", err.Error())
		}
		a.quit <- os.Interrupt
	}()
	<-a.quit
	a.Stop()
}

func (a *App) Stop() {
	log.Println("[graseful stop] stopping...")
	if a.router != nil {
		a.router.Shutdown()
	}
	if a.pool != nil {
		a.pool.Close()
	}
	close(a.quit)
}
