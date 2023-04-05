package app

import (
	"context"
	"fmt"
	"log"
	"nbrates/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newStorage(conf *config.AppConf) *pgxpool.Pool {
	ctx := context.Background()
	log.Println("[postgres-pool] init...")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		conf.Storage.UserName,
		conf.Storage.Password,
		conf.Storage.Host,
		conf.Storage.Port,
		conf.Storage.Name,
	)

	pg, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("[postgres-pool] init error: %s", err)
	}

	log.Println("[postgres-pool] check conn")

	conn, err := pg.Acquire(ctx)
	if err != nil {
		log.Fatalf("[postgres-pool] check conn error: %s", err)
	}

	conn.Release()
	log.Println("[postgres-pool] check conn OK")
	log.Println("[postgres-pool] init done")

	return pg
}
