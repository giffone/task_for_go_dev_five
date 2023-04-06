package service

import (
	"context"
	"nbrates/internal/domain"
	"time"
)

type Storage interface {
	Add(date time.Time, items []domain.Item)
}

type RestyClient interface {
	Do(ctx context.Context, date string) ([]domain.Item, error)
}

type Service interface {
	Add(ctx context.Context, date string) error
	Get(date, code string) error
}

func New(storage Storage, cli RestyClient) Service {
	return &service{
		storage: storage,
		cli:     cli,
	}
}

type service struct {
	storage Storage
	cli     RestyClient
}

func (s *service) Add(ctx context.Context, date string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	t, err := time.Parse("02.01.2006", date)
	if err != nil {
		return err
	}

	items, err := s.cli.Do(ctx, date)
	if err != nil {
		return err
	}

	go s.storage.Add(t, items)
	return nil
}

func (s *service) Get(date, code string) error {
	return nil
}
