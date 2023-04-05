package app

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

func newCli(baseUrl string) *resty.Client {
	r := resty.New()
	r.SetHeader("Content-Type", "application/xml;charset=UTF-8")
	r.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
		MaxVersion:         tls.VersionTLS12,
	})
	if baseUrl != "" {
		r.SetBaseURL(baseUrl)
	}
	return r
}
