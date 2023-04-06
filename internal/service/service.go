package service

import (
	"context"
	"log"
	"nbrates/internal/domain"
	"strconv"
	"time"
)

type Storage interface {
	Add(items []domain.ItemDTO)
	GetByDate(ctx context.Context, date time.Time) ([]domain.ItemDTO, error)
	GetByDateCode(ctx context.Context, date time.Time, code string) ([]domain.ItemDTO, error)
}

type RestyClient interface {
	Do(ctx context.Context, date string) ([]domain.Item, error)
}

type Service interface {
	Add(ctx context.Context, date string) error
	Get(ctx context.Context, date, code string) ([]domain.ItemDTO, error)
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
	// check date format
	t, err := time.Parse("02.01.2006", date)
	if err != nil {
		return err
	}
	// get data from National Bank api
	items, err := s.cli.Do(ctx, date)
	if err != nil {
		return err
	}
	// prepare data for db
	itemsDTO, i := dto(items, t)
	if i == 0 {
		return domain.NoRates
	}
	// add data to db
	go s.storage.Add(itemsDTO)
	return nil
}

func (s *service) Get(ctx context.Context, date, code string) ([]domain.ItemDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	// check date format
	t, err := time.Parse("02.01.2006", date)
	if err != nil {
		return nil, err
	}
	// get data from db
	if code != "" {
		return s.storage.GetByDateCode(ctx, t, code)
	}
	return s.storage.GetByDate(ctx, t)
}

func dto(items []domain.Item, date time.Time) ([]domain.ItemDTO, int) {
	if len(items) == 0 {
		return nil, 0
	}

	itemsDTO := make([]domain.ItemDTO, len(items))
	var i int

	for _, v := range items {
		value, err := strconv.ParseFloat(v.Description, 64)
		if err != nil {
			log.Printf("dto: items: convert rate: %s\n", err.Error())
			continue
		}
		itemsDTO[i] = domain.ItemDTO{
			Title: v.Fullname,
			Code:  v.Title,
			Value: value,
			Date:  date,
		}
		i++
	}
	// cut slice
	if i != len(items) {
		itemsDTO = itemsDTO[:i]
	}
	return itemsDTO, i
}
