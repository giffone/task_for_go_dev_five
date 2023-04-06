package app

import (
	"context"
	"fmt"
	"log"
	"nbrates/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newStorage(conf *config.AppConf) *pgxpool.Pool {
	// return nil
	ctx := context.Background()
	log.Println("[postgres-pool] init...")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Storage.UserName,
		conf.Storage.Password,
		conf.Storage.Host,
		conf.Storage.Port,
		conf.Storage.Name,
	)

	// connStr = "user=postgres password=postgres port=5432 dbname=postgres sslmode=disable"
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

	createTable(pg)

	return pg
}

func createTable(pool *pgxpool.Pool) {
	sqlQ := `CREATE TABLE IF NOT EXISTS r_currency (
		id SERIAL PRIMARY KEY,
		title VARCHAR(60),
		code VARCHAR(3),
		value NUMERIC(18,2),
		a_date DATE
	);`

	_, err := pool.Exec(context.Background(), sqlQ)
	if err != nil {
		log.Printf("create table: %s", err.Error())
		return
	}

	sqlUniq := "ALTER TABLE r_currency ADD CONSTRAINT unique_currency_date UNIQUE (code, a_date);"

	_, err = pool.Exec(context.Background(), sqlUniq)
	if err != nil {
		log.Printf("create table: %s", err.Error())
	}
}
