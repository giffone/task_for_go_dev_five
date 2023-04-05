package app

import (
	"fmt"
	"log"
	"nbrates/internal/api"
	"nbrates/internal/config"
	"nbrates/internal/service"

	"nbrates/internal/repository/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	conf   *config.AppConf
	router *fiber.App
	pool   *pgxpool.Pool
}

func New(conf *config.AppConf) *App {
	app := App{pool: newStorage(conf)}

	db := storage.New(app.pool)
	svc := service.New(db)
	hndl := api.New(svc)
	app.router = newRouter(hndl)
	return &app
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
