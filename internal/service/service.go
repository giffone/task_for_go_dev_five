package service

type Storage interface {
}

type RestyClient interface {
}

type Service interface {
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

func (s *service) Add() {

}

func (s *service) Get() {

}
