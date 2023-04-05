package app

import (
	"fmt"
	"log"
	"nbrates/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	conf   *config.AppConf
	router *fiber.App
	pool   *pgxpool.Pool
}

func New(conf *config.AppConf) *App {
	return &App{
		router: newRouter(),
		pool:   newStorage(conf),
	}
}

func (a *App) Start() error {
	addr := fmt.Sprintf("%s:%s", a.conf.Route.Host, a.conf.Route.Port)
	return a.router.Listen(addr)
}

func (a *App) Stop() {
	log.Println("[graseful stop] stopping...")
	if a.pool != nil {
		a.pool.Close()
	}

}
