package client

import (
	"nbrates/internal/service"

	"github.com/go-resty/resty/v2"
)

func New(cli *resty.Client) service.RestyClient {
	return &Cli{cli: cli}
}

type Cli struct {
	cli *resty.Client
}

func (c *Cli) Get() {

}
