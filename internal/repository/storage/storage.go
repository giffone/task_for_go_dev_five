package storage

import (
	"context"
	"log"
	"nbrates/internal/domain"
	"nbrates/internal/service"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var inserQuery = `INSERT INTO r_currency (title, code, value, a_date) VALUES ($1, $2, $3, $4);`

func New(pool *pgxpool.Pool) service.Storage {
	return &storage{pool: pool}
}

type storage struct {
	pool *pgxpool.Pool
}

func (s *storage) Add(date time.Time, items []domain.Item) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	batch := &pgx.Batch{}
	for _, item := range items {
		batch.Queue(
			inserQuery,
			item.Fullname,
			item.Title,
			item.Description,
			date,
		)
	}

	results := s.pool.SendBatch(ctx, batch)
	defer func() {
		if err := results.Close(); err != nil {
			log.Printf("close batch: %s\n", err)
		}
	}()

	for i := 0; i < len(items); i++ {
		_, err := results.Exec()
		if err != nil {
			log.Printf("exec batch: %s\n", err)
			break
		}
	}
}

func (s *storage) Get() {

}
