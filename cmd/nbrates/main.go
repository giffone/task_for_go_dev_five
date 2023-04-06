package main

import (
	"nbrates/internal/app"
	"nbrates/internal/config"
)

var confPath = "config/config.json"

func main() {
	conf := config.New(confPath)

	a := app.New(conf)
	a.Start()
	defer a.Stop()
}
