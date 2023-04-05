package config

import "errors"

type Storage struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (s Storage) validate() error {
	if s.Name == "" || s.Host == "" || s.Port == "" ||
		s.UserName == "" || s.Password == "" {
		return errors.New("empty storage config")
	}
	return nil
}
