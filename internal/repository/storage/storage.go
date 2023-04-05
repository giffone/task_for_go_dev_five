package storage

import (
	"nbrates/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(pool *pgxpool.Pool) service.Storage {
	return &storage{pool: pool}
}

type storage struct {
	pool *pgxpool.Pool
}

func (s *storage) Add() {

}

func(s *storage) Get() {
	
}