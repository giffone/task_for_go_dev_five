package config

import "errors"

type Route struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (r Route) validate() error {
	if r.Host == "" || r.Port == "" {
		return errors.New("empty route config")
	}
	return nil
}
