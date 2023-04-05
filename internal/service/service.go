package service

type Storage interface {
}

type Service interface {
}

func New(storage Storage) Service {
	return &service{storage: storage}
}

type service struct {
	storage Storage
}

func (s *service) Add() {

}

func (s *service) Get() {

}
