package storage

import (
	"context"
	"fmt"
	"log"
	"nbrates/internal/domain"
	"nbrates/internal/service"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var inserQuery = `INSERT INTO r_currency (title, code, value, a_date) VALUES ($1, $2, $3, $4);`
var getQuery = `SELECT (title, code, value, a_date) FROM r_currency WHERE a_date::date = $1;`
var getQuery2 = `SELECT (title, code, value, a_date) FROM r_currency WHERE a_date::date = $1 AND code = $2;`

func New(pool *pgxpool.Pool) service.Storage {
	return &storage{pool: pool}
}

type storage struct {
	pool *pgxpool.Pool
}

func (s *storage) Add(items []domain.ItemDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()

	batch := &pgx.Batch{}
	for _, item := range items {
		batch.Queue(
			inserQuery,
			item.Title,
			item.Code,
			item.Value,
			item.Date,
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

func (s *storage) Get(ctx context.Context, date time.Time, code string) error {
	if code == "" {
		rows, err := s.pool.Query(ctx, getQuery, date.Format("02.01.2006"))
		if err != nil {
			return fmt.Errorf("pgx: query: %w", err)
		}
		s.iterateRows(rows)
		return nil
	}
	return nil
}

func (s *storage) iterateRows(rows pgx.Rows) {
	for rows.Next() {

	}
}
