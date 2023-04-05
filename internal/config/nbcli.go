package config

import "errors"

type Nb struct {
	Link    string `json:"link"`
	Timeout string `json:"timeout"`
}

func (n Nb) validate() error {
	if n.Link == "" || n.Timeout == "" {
		return errors.New("empty route config")
	}
	return nil
}
