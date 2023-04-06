package client

import (
	"context"
	"encoding/xml"
	"fmt"
	"nbrates/internal/domain"
	"nbrates/internal/service"

	"github.com/go-resty/resty/v2"
)

var byDate = "?fdate="

func New(cli *resty.Client) service.RestyClient {
	return &Cli{cli: cli}
}

type Cli struct {
	cli *resty.Client
}

func (c *Cli) Do(ctx context.Context, date string) ([]domain.Item, error) {
	response, err := c.cli.R().
		SetContext(ctx).
		Get(byDate + date)
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, fmt.Errorf("resty: %s", response.String())
	}

	var result domain.Rates
	if err := xml.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	return result.Item, nil
}
