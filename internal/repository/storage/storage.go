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

var (
	inserQuery = `INSERT INTO r_currency (title, code, value, a_date)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (code, a_date) DO UPDATE SET
    title = EXCLUDED.title,
    value = EXCLUDED.value;`

	getQuery = `SELECT title, code, value, a_date
	FROM r_currency 
	WHERE a_date = $1;`
	
	getQuery2 = `SELECT title, code, value, a_date
	FROM r_currency 
	WHERE a_date = $1 
	AND code = $2;`
)

func New(pool *pgxpool.Pool) service.Storage {
	return &storage{pool: pool}
}

type storage struct {
	pool *pgxpool.Pool
}

func (s *storage) Add(items []domain.ItemDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
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

func (s *storage) GetByDate(ctx context.Context, date time.Time) ([]domain.ItemDTO, error) {
	rows, err := s.pool.Query(ctx, getQuery, date.Format("02.01.2006"))
	if err != nil {
		return nil, fmt.Errorf("pgx: query: %w", err)
	}
	defer rows.Close()
	return s.iterateRows(rows)
}

func (s *storage) GetByDateCode(ctx context.Context, date time.Time, code string) ([]domain.ItemDTO, error) {
	rows, err := s.pool.Query(ctx, getQuery2, date.Format("2006-01-02"), code)
	if err != nil {
		return nil, fmt.Errorf("pgx: query2: %w", err)
	}
	defer rows.Close()
	return s.iterateRows(rows)
}

func (s *storage) iterateRows(rows pgx.Rows) ([]domain.ItemDTO, error) {
	items := make([]domain.ItemDTO, 0, 50)
	for rows.Next() {
		item := domain.ItemDTO{}
		if err := rows.Scan(
			&item.Title,
			&item.Code,
			&item.Value,
			&item.Date,
		); err != nil {
			return nil, fmt.Errorf("iterate rows: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}
